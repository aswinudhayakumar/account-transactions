package signal

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	terminateSignals = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	globalSignal     signalWrap
	once             sync.Once
)

func init() {
	initializeGlobalSignal()
}

func initializeGlobalSignal() {
	once.Do(func() {
		globalSignal = signalWrap{
			sigCh:      make(chan os.Signal, 1),
			shutdownCh: make(chan struct{}),
		}
	})
}

type syncList struct {
	mu    sync.RWMutex
	funcs []func()
}

type signalWrap struct {
	funcs        syncList
	sigCh        chan os.Signal
	shutdownCh   chan struct{}
	shutdownOnce sync.Once
}

func (s *syncList) add(fn func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.funcs = append(s.funcs, fn)
}

func (s *syncList) list() []func() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.funcs
}

// Add registers a function to be executed during shutdown.
func Add(fn func()) {
	globalSignal.funcs.add(fn)
}

// New creates a channel that signals termination.
func New(sigs ...os.Signal) <-chan struct{} {
	return NewWithContext(context.Background(), sigs...).Done()
}

// NewWithContext creates a context that cancels on termination signals.
func NewWithContext(parent context.Context, sigs ...os.Signal) context.Context {
	if len(sigs) == 0 {
		sigs = terminateSignals
	}
	signal.Notify(globalSignal.sigCh, sigs...)

	ctx, cancel := context.WithCancel(parent)
	go func() {
		select {
		case <-parent.Done():
		case <-globalSignal.sigCh:
		case <-globalSignal.shutdownCh:
		}

		for _, fn := range globalSignal.funcs.list() {
			fn()
		}
		cancel()
	}()

	return ctx
}

// Shutdown triggers the global shutdown sequence.
func Shutdown() {
	globalSignal.shutdownOnce.Do(func() {
		close(globalSignal.shutdownCh)
	})
}

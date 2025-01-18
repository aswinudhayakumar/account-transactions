package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	"github.com/aswinudhayakumar/account-transactions/internal/migrator"
	"github.com/aswinudhayakumar/account-transactions/internal/signal"
	"go.uber.org/zap"
)

type env struct {
	AppPort string `envconfig:"APP_PORT"`

	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
	DBName     string `envconfig:"DB_NAME"`
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     string `envconfig:"DB_PORT"`
	SSLMode    string `envconfig:"SSL_MODE"`

	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"5s"`
}

func main() {
	// Initialize logger
	if err := logger.InitLogger(); err != nil {
		log.Fatalf("🛑 can't initialize zap logger: %v", err)
		return
	}
	defer logger.SyncLogger()

	// Initialise signal handling
	ctx := signal.NewWithContext(context.Background())

	// Initialise config from env variables
	var conf env
	failOnError(UnmarshalEnv(&conf), "🛑 failed to unmarshal env variables")

	// Initialise Database connection
	db, err := NewDBConnection(ctx, conf)
	failOnError(err, "🛑 failed to connect to database")
	defer db.Close()

	// Run Database migrations
	err = migrator.RunMigrations(db.DB)
	failOnError(err, "🛑 failed to apply db migrations")

	// Initialise the HTTP web server
	webServerConfig := buildWebServerConfig(conf, db)
	webServer := webServerConfig.InitWebServer()

	// Run the HTTP web server
	go func() {
		log.Println(fmt.Sprintf("🟢 HTTP server listening on port: %s 🚀", conf.AppPort))
		if err := webServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			failOnError(err, "🛑 failed to start HTTP server")
		}
	}()

	// Graceful shutdown
	signal.Add(func() {
		_, shutdownCancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout)
		defer shutdownCancel()

		if err := webServer.Shutdown(ctx); err != nil {
			log.Printf("🛑 Failed to shut down HTTP server: %v", err)
		}
		db.Close()
	})

	<-ctx.Done()
	log.Println("💻 Shutting down the application...")
}

func failOnError(err error, msg string) {
	if err != nil {
		logger.Log.Fatal(msg, zap.Error(err))
	}
}

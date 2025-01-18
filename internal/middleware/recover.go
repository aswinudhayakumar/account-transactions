package middleware

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	"go.uber.org/zap"
)

// RecoverInterceptor is a middleware to recover from panics in HTTP handlers
func RecoverInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rcv := recover(); rcv != nil {
				logger.Log.Warn("Recovered from panic", zap.Any("recovered_from", rcv))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

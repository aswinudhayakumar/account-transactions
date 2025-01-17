package main

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/internal/middleware"
	"github.com/go-chi/chi/v5"
)

// WebServerConfig holds the required configs for the HTTP web server
type WebServerConfig struct {
	conf env
}

// buildWebServerConfig builds and returns a new WebServerConfig
func buildWebServerConfig(conf env) WebServerConfig {
	return WebServerConfig{
		conf: conf,
	}
}

// InitWebServer initialises and returns a HTTP web server
func (ws *WebServerConfig) InitWebServer() *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.RecoverInterceptor)

	return &http.Server{
		Addr:    getServerPort(ws.conf.AppPort),
		Handler: r,
	}
}

func getServerPort(port string) string {
	return ":" + port
}

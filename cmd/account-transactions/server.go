package main

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/internal/middleware"
	accHandler "github.com/aswinudhayakumar/account-transactions/pkg/handler/accounts"
	trxHandler "github.com/aswinudhayakumar/account-transactions/pkg/handler/transactions"
	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

// WebServerConfig holds the required configs for the HTTP web server
type WebServerConfig struct {
	conf            env
	accountsHandler accHandler.AccountsHandler
	trxHandler      trxHandler.TransactionsHandler
}

// buildWebServerConfig builds and returns a new WebServerConfig
func buildWebServerConfig(conf env, db *sqlx.DB) WebServerConfig {
	// Initialise data repository
	dataRepo := repository.NewDataRepo(db)

	return WebServerConfig{
		conf:            conf,
		accountsHandler: accHandler.NewAccountsHandler(dataRepo),
		trxHandler:      trxHandler.NewTransactionsHandler(dataRepo),
	}
}

// InitWebServer initialises and returns a HTTP web server
func (ws *WebServerConfig) InitWebServer() *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.RecoverInterceptor)

	r.Route("/app/v1", func(r chi.Router) {
		// accounts API handlers
		r.Route("/accounts", func(r chi.Router) {
			r.Post("/", ws.accountsHandler.CreateAccount)
			r.Get("/{id}", ws.accountsHandler.GetAccountByAccountID)
		})

		// transactions API handlers
		r.Post("/transactions", ws.trxHandler.CreateTransaction)

	})

	return &http.Server{
		Addr:    getServerPort(ws.conf.AppPort),
		Handler: r,
	}
}

func getServerPort(port string) string {
	return ":" + port
}

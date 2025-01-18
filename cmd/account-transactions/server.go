package main

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/internal/middleware"
	writer "github.com/aswinudhayakumar/account-transactions/internal/writter"
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
func InitWebServer(config WebServerConfig) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.RecoverInterceptor)

	r.Route("/app/v1", func(r chi.Router) {
		// test API
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			writer.WriteJSON(w, http.StatusOK, map[string]string{"msg": "helloworld!"})
		})

		// accounts API handlers
		r.Route("/accounts", func(r chi.Router) {
			r.Post("/", config.accountsHandler.CreateAccount)
		})
		r.Route("/accounts/{id}", func(r chi.Router) {
			r.Get("/", config.accountsHandler.GetAccountByAccountID)
		})

		// transactions API handlers
		r.Route("/transactions", func(r chi.Router) {
			r.Post("/", config.trxHandler.CreateTransaction)
		})
	})

	return &http.Server{
		Addr:    getServerPort(config.conf.AppPort),
		Handler: r,
	}
}

func getServerPort(port string) string {
	return ":" + port
}

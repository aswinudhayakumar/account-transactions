package handler

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
)

// AccountsHandler is an interface for handling accounts-related API requests.
type AccountsHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccountByAccountID(w http.ResponseWriter, r *http.Request)
}

// accountsHandler object.
type accountsHandler struct {
	DataRepo repository.DataRepo
}

// NewAccountsHandler initializes and returns a new AccountsHandler.
func NewAccountsHandler(dataRepo repository.DataRepo) AccountsHandler {
	return &accountsHandler{
		DataRepo: dataRepo,
	}
}

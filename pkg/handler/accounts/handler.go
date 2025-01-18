package handler

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
)

type AccountsHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccountByAccountID(w http.ResponseWriter, r *http.Request)
}

type accountsHandler struct {
	DataRepo repository.DataRepo
}

func NewAccountsHandler(dataRepo repository.DataRepo) AccountsHandler {
	return &accountsHandler{
		DataRepo: dataRepo,
	}
}

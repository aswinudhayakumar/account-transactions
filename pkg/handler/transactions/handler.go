package handler

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
)

// TransactionsHandler is an interface for handling transactions-related API requests.
type TransactionsHandler interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

// transactionsHandler object.
type transactionsHandler struct {
	DataRepo repository.DataRepo
}

// NewTransactionsHandler initializes and returns a new TransactionsHandler.
func NewTransactionsHandler(dataRepo repository.DataRepo) TransactionsHandler {
	return &transactionsHandler{
		DataRepo: dataRepo,
	}
}

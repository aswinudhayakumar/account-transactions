package handler

import (
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
)

type TransactionsHandler interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

type transactionsHandler struct {
	DataRepo repository.DataRepo
}

func NewTransactionsHandler(dataRepo repository.DataRepo) TransactionsHandler {
	return &transactionsHandler{
		DataRepo: dataRepo,
	}
}

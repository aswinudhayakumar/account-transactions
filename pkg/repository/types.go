package repository

import (
	"errors"
	"time"
)

var (
	ErrAccountIDNotExists       = errors.New("Account ID not exists")
	ErrOperationTypeIDNotExists = errors.New("Operation Type ID not exists")
)

type CreateAccountReqParams struct {
	DocumentNumber string `db:"document_number"`
}

type AccountResponse struct {
	AccountID      int       `db:"account_id"`
	DocumentNumber string    `db:"document_number"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type CreateTransactionReqParams struct {
	AccountID       int
	OperationTypeID int
	Amount          float64
}

type validateCreateTrx struct {
	IsAccountExists         bool `db:"is_account_exists"`
	IsOperationTypeIDExists bool `db:"is_operation_type_id_exists"`
}

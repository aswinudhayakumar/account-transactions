package repository

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrAccountIDNotExists       = errors.New("account id not exists")
	ErrOperationTypeIDNotExists = errors.New("operation type id not exists")
)

// CreateAccountReqParams is the request object for CreateAccount method.
type CreateAccountReqParams struct {
	DocumentNumber string `db:"document_number"`
}

// AccountResponse is the response object which holds Account data.
type AccountResponse struct {
	AccountID      int       `db:"account_id"`
	DocumentNumber string    `db:"document_number"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// CreateTransactionReqParams is the request object for CreateTransaction method.
type CreateTransactionReqParams struct {
	AccountID       int
	OperationTypeID int
	Amount          float64
	Balance         float64
}

type validateCreateTrx struct {
	IsAccountExists         bool `db:"is_account_exists"`
	IsOperationTypeIDExists bool `db:"is_operation_type_id_exists"`
}

type getOperationTypeResp struct {
	Transaction_type string `db:"transaction_type"`
}

type GetNegativeTransactionsResp struct {
	TransactionID int             `db:"transaction_id"`
	Balance       sql.NullFloat64 `db:"balance"`
}

type UpdateTransactionBalances struct {
	TransactionID int     `db:"transaction_id"`
	Balance       float64 `db:"balance"`
}

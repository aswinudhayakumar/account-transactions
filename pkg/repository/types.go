package repository

import (
	"time"
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

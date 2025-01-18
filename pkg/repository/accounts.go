package repository

import (
	"context"
)

const (
	createAccountQuery = `INSERT INTO accounts (document_number) VALUES ($1);`

	getAccountByAccountIDQuery = `SELECT account_id, document_number, created_at, updated_at FROM accounts WHERE account_id=$1;`
)

// CreateAccount creates a new Account.
func (dr *dataRepo) CreateAccount(ctx context.Context, req CreateAccountReqParams) error {
	_, err := dr.db.ExecContext(
		ctx,
		createAccountQuery,
		req.DocumentNumber,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetAccountByAccountID returns the Account from the database using the provided account ID.
func (dr *dataRepo) GetAccountByAccountID(ctx context.Context, accountID int) (*AccountResponse, error) {
	var res AccountResponse
	err := dr.db.GetContext(
		ctx,
		&res,
		getAccountByAccountIDQuery,
		accountID,
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

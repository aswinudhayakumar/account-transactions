package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	createAccountQuery = `INSERT INTO accounts (document_number) VALUES ($1);`

	getAccountByAccountIDQuery = `SELECT account_id, document_number, created_at, updated_at FROM accounts WHERE account_id=$1;`
)

func (dr *dataRepo) CreateAccount(ctx context.Context, req CreateAccountReqParams) error {
	return dr.execTxn(ctx, func(tx *sqlx.Tx) error {
		_, err := tx.ExecContext(
			ctx,
			createAccountQuery,
			req.DocumentNumber,
		)
		if err != nil {
			return err
		}

		return nil
	})
}

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

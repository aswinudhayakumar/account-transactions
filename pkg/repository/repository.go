//go:generate mockery --name=DataRepo --output=../../internal/mocks
package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DataRepo interface {
	CreateAccount(ctx context.Context, req CreateAccountReqParams) error
	GetAccountByAccountID(ctx context.Context, accountID int) (*AccountResponse, error)
	CreateTransaction(ctx context.Context, req CreateTransactionReqParams) error
}

type dataRepo struct {
	db *sqlx.DB
}

func NewDataRepo(db *sqlx.DB) DataRepo {
	return &dataRepo{
		db: db,
	}
}

func (dr *dataRepo) execTxn(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := dr.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin a transaction: %w", err)
	}

	err = fn(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("failed to rollback the transaction: %w", rbErr)
		}
		return fmt.Errorf("failed to complete the db transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit the transaction: %w", err)
	}
	return nil
}

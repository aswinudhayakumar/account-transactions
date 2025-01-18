//go:generate mockery --name=DataRepo --output=../../internal/mocks
package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// DataRepo is an interface which provides methods for database related operations.
type DataRepo interface {
	CreateAccount(ctx context.Context, req CreateAccountReqParams) error
	GetAccountByAccountID(ctx context.Context, accountID int) (*AccountResponse, error)
}

// dataRepo object.
type dataRepo struct {
	db *sqlx.DB
}

// NewDataRepo initializes and returns a new DataRepo.
func NewDataRepo(db *sqlx.DB) DataRepo {
	return &dataRepo{
		db: db,
	}
}

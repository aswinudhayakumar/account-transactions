package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	createTransactionQuery = `INSERT INTO transactions (account_id, operation_type_id, amount) VALUES ($1, $2, $3);`

	validateCreateTrxQuery = `
	SELECT 
		EXISTS (SELECT 1 FROM accounts WHERE account_id=$1) AS is_account_exists,
		EXISTS (SELECT 1 FROM operations_types WHERE operation_type_id=$2) AS is_operation_type_id_exists;
	`
)

func (dr *dataRepo) CreateTransaction(ctx context.Context, req CreateTransactionReqParams) error {
	return dr.execTxn(ctx, func(tx *sqlx.Tx) error {
		// Validate if both accountID and operationTypeID exists before creating a transaction
		var validation validateCreateTrx
		if err := tx.GetContext(
			ctx,
			&validation,
			validateCreateTrxQuery,
			req.AccountID,
			req.OperationTypeID,
		); err != nil {
			return err
		}

		if !validation.IsAccountExists {
			return ErrAccountIDNotExists
		}

		if !validation.IsOperationTypeIDExists {
			return ErrOperationTypeIDNotExists
		}

		// TODO: Get operation_type from operation_type_id and
		// update the amount (postiive or negative) based on it

		// Create a transaction
		_, err := tx.ExecContext(
			ctx,
			createTransactionQuery,
			req.AccountID,
			req.OperationTypeID,
			req.Amount,
		)
		if err != nil {
			return err
		}

		return nil
	})
}

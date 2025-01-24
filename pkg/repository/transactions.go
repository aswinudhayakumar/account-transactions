package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	createTransactionQuery = `INSERT INTO transactions (account_id, operation_type_id, amount, balance) VALUES ($1, $2, $3, $4);`

	validateCreateTrxQuery = `
	SELECT 
		EXISTS (SELECT 1 FROM accounts WHERE account_id=$1) AS is_account_exists,
		EXISTS (SELECT 1 FROM operations_types WHERE operation_type_id=$2) AS is_operation_type_id_exists;
	`

	getOperationTypeQuery = `SELECT transaction_type from operations_types where operation_type_id=$1;`

	getNegativeTransactionsQuery = `select transaction_id, balance from transactions where balance < 0 and account_id=$1;`

	updateTransaction = `update transactions set balance=$1 where transaction_id=$2;`
)

// CreateTransaction creates a new transaction.
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
			req.Balance,
		)
		if err != nil {
			return err
		}

		return nil
	})
}

func (dr *dataRepo) GetOperationType(ctx context.Context, operation_type_id int) (string, error) {
	var res getOperationTypeResp
	err := dr.db.GetContext(
		ctx,
		&res,
		getOperationTypeQuery,
		operation_type_id,
	)
	if err != nil {
		return "", err
	}

	return res.Transaction_type, nil
}

func (dr *dataRepo) GetNegativeTransactions(ctx context.Context, accountID int) ([]GetNegativeTransactionsResp, error) {
	var resp []GetNegativeTransactionsResp
	err := dr.db.SelectContext(
		ctx,
		&resp,
		getNegativeTransactionsQuery,
		accountID,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (dr *dataRepo) UpdateOldTransactionBalance(ctx context.Context, req UpdateTransactionBalances) error {
	return dr.execTxn(
		ctx, func(tx *sqlx.Tx) error {
			_, err := dr.db.ExecContext(
				ctx,
				updateTransaction,
				req.Balance,
				req.TransactionID,
			)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

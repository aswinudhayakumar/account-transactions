package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	"github.com/aswinudhayakumar/account-transactions/internal/writer"
	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
	"go.uber.org/zap"
)

// CreateTransaction handles the creation of a new transaction.
func (h *transactionsHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Decode the request params
	var req CreateTrxReqParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Failed to decode HTTP request payload", zap.Error(err))
		writer.WriteJSONError(
			w,
			http.StatusBadRequest,
			writer.ErrorDescription{
				Title:  writer.ErrTitleInvalidRequestPayload,
				Code:   writer.ErrCodeInvalidRequest,
				Detail: "The request payload is malformed or invalid.",
			},
		)
		return
	}

	trxType, err := h.DataRepo.GetOperationType(r.Context(), req.OperationTypeID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			writer.WriteJSONError(
				w,
				http.StatusNotFound,
				writer.ErrorDescription{
					Title:  writer.ErrTitleInvalidRequestPayload,
					Code:   writer.ErrCodeDataNotFound,
					Detail: err.Error(),
				},
			)
			return
		}

		writer.WriteJSONError(
			w,
			http.StatusBadRequest,
			writer.ErrorDescription{
				Title:  writer.ErrTitleInvalidRequestPayload,
				Code:   writer.ErrCodeUnexpectedError,
				Detail: err.Error(),
			},
		)
		return
	}

	if trxType == "debit" {
		req.Amount = -req.Amount
	} else {
		transactions, err := h.DataRepo.GetNegativeTransactions(r.Context(), req.AccountID)
		if err != nil {
			writer.WriteJSONError(
				w,
				http.StatusBadRequest,
				writer.ErrorDescription{
					Title:  writer.ErrTitleInvalidRequestPayload,
					Code:   writer.ErrCodeUnexpectedError,
					Detail: err.Error(),
				},
			)
			return
		}

		logger.Log.Info("before --> ", zap.Any("transactions", transactions))

		amount := req.Amount

		var balance float64
		var processed []repository.UpdateTransactionBalances

		for _, trx := range transactions {
			if amount > 0 {
				if trx.Balance.Valid {
					balance = -trx.Balance.Float64
				}
				updatedAmount := amount - balance
				if updatedAmount > 0 {
					processed = append(processed, repository.UpdateTransactionBalances{
						TransactionID: trx.TransactionID,
						Balance:       0,
					})
				} else {
					updatedAmount = 0
				}

				amount = updatedAmount
			}
		}

		logger.Log.Info("after --> ", zap.Any("processed", processed))
		fmt.Println("final Amoutn -> ", req.Amount, amount)

	}

	// Create a transaction
	err = h.DataRepo.CreateTransaction(
		r.Context(),
		repository.CreateTransactionReqParams{
			AccountID:       req.AccountID,
			OperationTypeID: req.OperationTypeID,
			Amount:          req.Amount,
			Balance:         req.Amount,
		})
	if err != nil {
		// Handle errors
		if errors.Is(err, repository.ErrAccountIDNotExists) || errors.Is(err, repository.ErrOperationTypeIDNotExists) {
			logger.Log.Error("Failed to create transaction", zap.Error(err))
			writer.WriteJSONError(
				w,
				http.StatusBadRequest,
				writer.ErrorDescription{
					Title:  writer.ErrTitleInvalidRequestPayload,
					Code:   writer.ErrCodeInvalidRequest,
					Detail: err.Error(),
				},
			)
			return
		}

		logger.Log.Error("Database call failed for CreateTransaction request", zap.Error(err))
		writer.WriteJSONError(
			w,
			http.StatusInternalServerError,
			writer.ErrorDescription{
				Title:  writer.ErrTitleUnexpectedError,
				Code:   writer.ErrCodeUnexpectedError,
				Detail: err.Error(),
			},
		)
		return
	}

	// Send success response
	resp := CreateTrxResponse{
		Status: StatusSuccess,
	}
	if err := writer.WriteJSON(w, http.StatusCreated, resp); err != nil {
		logger.Log.Error("Error writting success response for CreateTransaction request", zap.Error(err))
		writer.WriteJSONError(
			w,
			http.StatusInternalServerError,
			writer.ErrorDescription{
				Title:  writer.ErrTitleUnexpectedError,
				Code:   writer.ErrCodeUnexpectedError,
				Detail: err.Error(),
			},
		)
		return
	}
}

// Add valiadtion if requried
// func validateCreateTransaction(req CreateTrxReqParams) *validator.ValidationErrors {
// 	errors := validator.NewValidationErrors()

// 	if len(errors.Errors) > 0 {
// 		return errors
// 	}

// 	return nil
// }

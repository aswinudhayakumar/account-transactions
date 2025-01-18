package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	writer "github.com/aswinudhayakumar/account-transactions/internal/writter"
	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
	"go.uber.org/zap"
)

func (h *transactionsHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	// Decode the request params
	var req CreateTrxReqParams
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Log.Error("Failed to decode HTTP request payload", zap.Error(err))
		writer.WriteJSONError(
			w,
			http.StatusBadRequest,
			writer.ErrorDescription{
				Title:  "Invalid Request Payload",
				Code:   "invalid_request",
				Detail: "The request payload is malformed or invalid.",
			},
		)
		return
	}

	// Create a transaction
	err := h.DataRepo.CreateTransaction(
		r.Context(),
		repository.CreateTransactionReqParams{
			AccountID:       req.AccountID,
			OperationTypeID: req.OperationTypeID,
			Amount:          req.Amount,
		})
	if err != nil {
		// Handle errors
		if errors.Is(err, repository.ErrAccountIDNotExists) || errors.Is(err, repository.ErrOperationTypeIDNotExists) {
			logger.Log.Error("Failed to create transaction", zap.Error(err))
			writer.WriteJSONError(
				w,
				http.StatusBadRequest,
				writer.ErrorDescription{
					Title:  "Invalid Request Payload",
					Code:   "invalid_request",
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
				Title:  "Unexpected Error",
				Code:   "unexpected_error",
				Detail: err.Error(),
			},
		)
		return
	}

	// Send success response
	resp := CreateTrxResponse{
		Status: true,
	}
	if err := writer.WriteJSON(w, http.StatusCreated, resp); err != nil {
		logger.Log.Error("Error writting success response for CreateTransaction request", zap.Error(err))
		writer.WriteJSONError(
			w,
			http.StatusInternalServerError,
			writer.ErrorDescription{
				Title:  "Unexpected Error",
				Code:   "unexpected_error",
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

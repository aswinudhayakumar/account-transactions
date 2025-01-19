package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	"github.com/aswinudhayakumar/account-transactions/internal/validator"
	"github.com/aswinudhayakumar/account-transactions/internal/writer"
	"github.com/aswinudhayakumar/account-transactions/pkg/repository"
	"go.uber.org/zap"
)

// CreateAccount handles the creation of a new account.
func (h *accountsHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Decode the request params
	var req CreateAccountReqParams
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

	// Validate the request params
	if validationErrs := validateCreateAccountRequest(req); validationErrs != nil {
		logger.Log.Error("Validation failed for CreateAccount request", zap.String("validation_errors", validationErrs.Error()))
		writer.WriteJSONError(
			w,
			http.StatusBadRequest,
			writer.ErrorDescription{
				Title:  writer.ErrTilteValidationFailed,
				Code:   writer.ErrCodeInvalidRequest,
				Detail: validationErrs.Error(),
			},
		)
		return
	}

	// Create new account
	err := h.DataRepo.CreateAccount(
		r.Context(),
		repository.CreateAccountReqParams{
			DocumentNumber: req.DocumentNumber,
		},
	)
	if err != nil {
		logger.Log.Error("Database call failed for CreateAccount request", zap.Error(err))
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
	resp := CreateAccountResponse{
		Status: StatusSuccess,
	}
	if err := writer.WriteJSON(w, http.StatusCreated, resp); err != nil {
		logger.Log.Error("Error writting success response for CreateAccount request", zap.Error(err))
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

// validateCreateAccountRequest validates the request object for CreateAccount API handler.
func validateCreateAccountRequest(req CreateAccountReqParams) *validator.ValidationErrors {
	errors := validator.NewValidationErrors()

	if len(req.DocumentNumber) < 3 || len(req.DocumentNumber) > 255 {
		errors.Add("document_number", "Document number must be between 3 and 255 characters in length.")
	}

	if len(errors.Errors) > 0 {
		return errors
	}

	return nil
}

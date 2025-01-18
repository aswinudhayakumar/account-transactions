package handler

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aswinudhayakumar/account-transactions/internal/logger"
	"github.com/aswinudhayakumar/account-transactions/internal/writer"
	"go.uber.org/zap"
)

// GetAccountByAccountID retrieves an account using the Account ID.
func (h *accountsHandler) GetAccountByAccountID(w http.ResponseWriter, r *http.Request) {
	// path : "app/v1/accounts/{id}"
	path := r.URL.Path
	parts := strings.Split(path, "/")
	accountIDStr := parts[4]

	// Get accountID from request URL
	accountID, err := strconv.Atoi(accountIDStr)
	if err != nil {
		logger.Log.Error("Failed to get query param from request URL", zap.Error(err))
		writer.WriteJSONError(
			w,
			http.StatusBadRequest,
			writer.ErrorDescription{
				Title:  "Invalid Account ID",
				Code:   writer.ErrCodeInvalidRequest,
				Detail: err.Error(),
			},
		)
		return
	}

	// Get the account using account ID
	dbResp, err := h.DataRepo.GetAccountByAccountID(r.Context(), accountID)
	if err != nil {
		logger.Log.Error("Database call failed for GetAccountByAccountID request", zap.Error(err))
		// Return 404 error if data not found
		if errors.Is(err, sql.ErrNoRows) {
			writer.WriteJSONError(
				w,
				http.StatusNotFound,
				writer.ErrorDescription{
					Title:  writer.ErrTitleDataNotFound,
					Code:   writer.ErrCodeDataNotFound,
					Detail: err.Error(),
				},
			)
			return
		}

		// Return 500 internal server error for other errors
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
	resp := AccountResponse{
		AccountID:      dbResp.AccountID,
		DocumentNumber: dbResp.DocumentNumber,
		CreatedAt:      dbResp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      dbResp.UpdatedAt.Format(time.RFC3339),
	}
	if err := writer.WriteJSON(w, http.StatusOK, resp); err != nil {
		logger.Log.Error("Error writting success response for GetAccountByAccountID request", zap.Error(err))
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

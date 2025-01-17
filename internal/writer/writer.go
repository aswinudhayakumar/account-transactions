package writer

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
)

const (
	// error codes
	ErrCodeInvalidRequest  = "invalid_request"
	ErrCodeUnexpectedError = "unexpected_error"
	ErrCodeDataNotFound    = "data_not_found"

	// error titles
	ErrTitleInvalidRequestPayload = "Invalid Request Payload"
	ErrTilteValidationFailed      = "Validation Failed"
	ErrTitleUnexpectedError       = "Unexpected Error"
	ErrTitleDataNotFound          = "Requested Data Not Found"
)

// Errors
var (
	ErrEmptyHTTPStatus   = errors.New("HTTP status must be set")
	ErrEmptyErrorMessage = errors.New("error message cannot be empty")
)

type (
	FieldError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ErrorDescription struct {
		ID     string `json:"id"`
		Code   string `json:"code"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
		Status int    `json:"status"`
	}

	ErrorResponse struct {
		Errors []ErrorDescription `json:"errors"`
	}
)

// WriteJSON writes the provided data as JSON with a given status code.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	if status == 0 {
		return ErrEmptyHTTPStatus
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// WriteJSONError writes an error response in JSON format.
func WriteJSONError(w http.ResponseWriter, status int, errDesc ErrorDescription, sources ...FieldError) error {
	if errDesc.Title == "" {
		return ErrEmptyErrorMessage
	}

	errDesc.Status = status

	var errResps []ErrorDescription
	if len(sources) > 0 {
		for _, source := range sources {
			resp, err := configureErrorResponse(errDesc, &source)
			if err != nil {
				return err
			}
			errResps = append(errResps, resp)
		}
	} else {
		resp, err := configureErrorResponse(errDesc, nil)
		if err != nil {
			return err
		}
		errResps = append(errResps, resp)
	}
	return WriteJSON(w, status, ErrorResponse{Errors: errResps})
}

func configureErrorResponse(resp ErrorDescription, source *FieldError) (ErrorDescription, error) {
	if resp.ID == "" {
		id, err := uuid.NewV7()
		if err != nil {
			return ErrorDescription{}, fmt.Errorf("failed to generate UUID: %w", err)
		}
		resp.ID = id.String()
	}

	return resp, nil
}

package handler

// CreateAccountReqParams is the request object for CreateAccount API.
type CreateAccountReqParams struct {
	DocumentNumber string `json:"document_number"`
}

// CreateAccountResponse is the response object for CreateAccount API.
type CreateAccountResponse struct {
	Status bool `json:"status"`
}

// AccountResponse is the response object, which holds Account data.
type AccountResponse struct {
	AccountID      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

package handler

type CreateAccountReqParams struct {
	DocumentNumber string `json:"document_number"`
}

type CreateAccountResponse struct {
	Status bool `json:"status"`
}

type AccountResponse struct {
	AccountID      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

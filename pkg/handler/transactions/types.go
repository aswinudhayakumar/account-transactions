package handler

// CreateTrxReqParams is the request object for CreateTransaction API.
type CreateTrxReqParams struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

// CreateTrxResponse is the response object for CreateTransaction API.
type CreateTrxResponse struct {
	Status bool `json:"status"`
}

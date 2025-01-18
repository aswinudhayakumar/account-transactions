package handler

type CreateTrxReqParams struct {
	AccountID       int     `json:"account_id"`
	OperationTypeID int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type CreateTrxResponse struct {
	Status bool `json:"status"`
}

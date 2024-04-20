package dto

type CreateTransactionRequest struct {
	Title  string  `json:"title"`
	Note   *string `json:"note"`
	Amount uint64  `json:"amount"`
	Type   string  `json:"type"`
}

type UpdateTransactionRequest struct {
	Title  string  `json:"title"`
	Note   *string `json:"note"`
	Amount uint64  `json:"amount"`
	Type   string  `json:"type"`
}

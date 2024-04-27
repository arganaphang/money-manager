package dto

import "github.com/arganaphang/money-manager/pkg"

type CreateTransactionRequest struct {
	Title  string  `json:"title" form:"title"`
	Note   *string `json:"note" form:"form"`
	Amount uint64  `json:"amount" form:"amount"`
	Type   string  `json:"type" form:"type"`
}

type UpdateTransactionByIDRequest struct {
	Title  string  `json:"title" form:"title"`
	Note   *string `json:"note" form:"note"`
	Amount uint64  `json:"amount" form:"amount"`
	Type   string  `json:"type" form:"type"`
}

type GetTransactionsRequest struct {
	pkg.Pagination
}

package dto

import (
	"github.com/arganaphang/money-manager/internal/model"
	"github.com/arganaphang/money-manager/pkg"
)

type CreateTransactionResponse struct {
	Message string `json:"message"`
}

type GetTransactionsResponse struct {
	Message string              `json:"message"`
	Data    []model.Transaction `json:"data"`
	Meta    pkg.Pagination      `json:"meta"`
}

type GetTransactionByIDResponse struct {
	Message string             `json:"message"`
	Data    *model.Transaction `json:"data"`
}

type UpdateTransactionByIDResponse struct {
	Message string `json:"message"`
}

type DeleteTransactionByIDResponse struct {
	Message string `json:"message"`
}

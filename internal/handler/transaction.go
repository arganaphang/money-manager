package handler

import (
	"net/http"

	"github.com/arganaphang/money-manager/internal/dto"
	"github.com/arganaphang/money-manager/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITransactionHandler interface {
	CreateTransaction(ctx *gin.Context)
	GetTransactions(ctx *gin.Context)
	GetTransactionByID(ctx *gin.Context)
	UpdateTransactionByID(ctx *gin.Context)
	DeleteTransactionByID(ctx *gin.Context)
}

type TransactionHandler struct {
	Services service.Services
}

func NewTransactionHandler(services service.Services) ITransactionHandler {
	return &TransactionHandler{Services: services}
}

func (h TransactionHandler) CreateTransaction(ctx *gin.Context) {
	var data dto.CreateTransactionRequest
	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := h.Services.TransactionServices.CreateTransaction(ctx, data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.CreateTransactionResponse{
		Message: "transaction created",
	})
}

func (h TransactionHandler) GetTransactions(ctx *gin.Context) {
	var data dto.GetTransactionsRequest
	if err := ctx.ShouldBindQuery(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result, err := h.Services.TransactionServices.GetTransactions(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.GetTransactionsResponse{
		Message: "get transactions",
		Data:    result,
		Meta:    data.Pagination,
	})
}

func (h TransactionHandler) GetTransactionByID(ctx *gin.Context) {
	result, err := h.Services.TransactionServices.GetTransactionByID(ctx, uuid.New())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.GetTransactionByIDResponse{
		Message: "get transaction by id",
		Data:    result,
	})
}

func (h TransactionHandler) UpdateTransactionByID(ctx *gin.Context) {
	err := h.Services.TransactionServices.UpdateTransactionByID(ctx, uuid.New(), dto.UpdateTransactionByIDRequest{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.UpdateTransactionByIDResponse{
		Message: "transaction updated",
	})
}

func (h TransactionHandler) DeleteTransactionByID(ctx *gin.Context) {
	err := h.Services.TransactionServices.DeleteTransactionByID(ctx, uuid.New())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.DeleteTransactionByIDResponse{
		Message: "delete transaction by id",
	})
}

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
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "transaction created",
	})
}

func (h TransactionHandler) GetTransactions(ctx *gin.Context) {
	result, err := h.Services.TransactionServices.GetTransactions(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "get transactions",
		"data":    result,
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
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "get transaction by id",
		"data":    result,
	})
}

func (h TransactionHandler) UpdateTransactionByID(ctx *gin.Context) {
	err := h.Services.TransactionServices.UpdateTransactionByID(ctx, uuid.New(), dto.UpdateTransactionRequest{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "transaction updated",
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
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "delete transaction by id",
	})
}

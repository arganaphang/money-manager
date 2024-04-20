package main

import (
	"net/http"

	"github.com/arganaphang/money-manager/internal/handler"
	"github.com/arganaphang/money-manager/internal/repository"
	"github.com/arganaphang/money-manager/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/healthz", health)

	// Create DB Connection
	// Create Repository
	repositories := repository.Repositories{
		TransactionRepository: repository.NewTransactionRepository(),
	}
	// Create Service
	services := service.Services{
		TransactionServices: service.NewTransactionService(repositories),
	}
	// Create Handler
	trxHandler := handler.NewTransactionHandler(services)

	trx := app.Group("/transaction")
	trx.GET("", trxHandler.GetTransactions)
	trx.GET(":id", trxHandler.GetTransactionByID)
	trx.POST("", trxHandler.CreateTransaction)
	trx.PUT(":id", trxHandler.UpdateTransactionByID)
	trx.DELETE(":id", trxHandler.DeleteTransactionByID)

	app.Run("0.0.0.0:8000")
}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"message": "OK hehe",
	})
}

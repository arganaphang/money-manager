package main

import (
	"net/http"

	"github.com/arganaphang/money-manager/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.GET("/healthz", health)

	trx := app.Group("/transaction")
	trx.GET("", handler.GetTransactions)
	trx.GET(":id", handler.GetTransactionByID)
	trx.POST("", handler.CreateTransaction)
	trx.PUT(":id", handler.UpdateTransactionByID)
	trx.DELETE(":id", handler.DeleteTransactionByID)

	app.Run("0.0.0.0:8000")
}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"message": "OK hehe",
	})
}

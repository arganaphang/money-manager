package main

import (
	"log"
	"net/http"
	"os"

	"github.com/arganaphang/money-manager/internal/handler"
	"github.com/arganaphang/money-manager/internal/repository"
	"github.com/arganaphang/money-manager/internal/service"
	"github.com/gin-gonic/gin"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.Default().SetTrustedProxies(nil)
	app := gin.New()

	app.GET("/healthz", health)

	// Create DB Connection
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln(err)
	}
	// Create Repository
	repositories := repository.Repositories{
		TransactionRepository: repository.NewTransactionRepository(db),
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

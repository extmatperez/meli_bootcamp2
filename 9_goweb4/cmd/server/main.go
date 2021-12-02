package main

import (
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/cmd/server/handler"
	transaction "github.com/extmatperez/meli_bootcamp2/9_goweb4/internal/transaction"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	router := gin.Default()

	repo := transaction.NewRepository()
	service := transaction.NewService(repo)
	controller := handler.NewTransaction(service)

	transactionURL := router.Group("/transactions")

	transactionURL.GET("/", controller.GetAll())
	transactionURL.GET("/:id", controller.GetByID())
	transactionURL.GET("/receivers/:receiver", controller.GetByReceiver())
	transactionURL.POST("/", controller.Store())
	//transactionURL.POST("/", controller.CreateTransaction())
	transactionURL.PUT("/:id", controller.UpdateTransaction())
	transactionURL.PATCH("/:id/:amount", controller.UpdateAmount())
	transactionURL.DELETE("/:id", controller.DeleteTransaction())

	router.Run(":9090")
}

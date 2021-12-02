package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/turn_morning/exercise_1/cmd/server/handler"
	transaction "github.com/extmatperez/meli_bootcamp2/8_goweb3/turn_morning/exercise_1/internal/transaction"
	"github.com/gin-gonic/gin"
)

func main() {

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

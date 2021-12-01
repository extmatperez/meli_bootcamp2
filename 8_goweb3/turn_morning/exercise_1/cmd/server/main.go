package main

import (
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/turn_afternoon/cmd/server/handler"
	transaction "github.com/extmatperez/meli_bootcamp2/7_goweb2/turn_afternoon/internal/transaction"
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
	//transactionURL.PUT("/", controller.UpdateTransaction())
	//transactionURL.PATCH("/", controller.UpdateAmount())
	//transactionURL.DELETE("/", controller.DeleteTransaction())

	router.Run(":9090")
}

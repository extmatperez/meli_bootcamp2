package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/soto_jose/8_goweb3/cmd/server/handler"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/soto_jose/8_goweb3/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	controller := handler.NewTransaction(service)

	transactions := router.Group("/transactions")

	transactions.GET("/transactions", controller.GetAll())
	transactions.POST("/transactions", controller.Store())

	router.Run()
}

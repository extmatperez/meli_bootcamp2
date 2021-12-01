package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lean1097/meli_bootcamp2/7_goweb2/TT/internal"
	"github.com/lean1097/meli_bootcamp2/7_goweb2/TT/server/handler"
)

func main() {
	repo := internal.NewRepository()
	service := internal.NewService(repo)
	product := handler.NewProduct(service)

	router := gin.Default()
	products := router.Group("/products")

	products.POST("/", product.Store())
	products.GET("/", product.GetAll())

	router.Run()
}

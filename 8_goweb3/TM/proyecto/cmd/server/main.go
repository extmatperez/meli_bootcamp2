package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TM/proyecto/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/TM/proyecto/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService()
	handler := handler.NewProduct(service)

	router.GET("/products/get", handler.GetAll())
	router.POST("/products/add", handler.Save())
	router.PUT("/:id", handler.Update())

	router.Run()
}

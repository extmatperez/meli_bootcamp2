package main

import (
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/cmd/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService()
	handler := handler.NewProduct(service)

	router.GET("/products/get", handler.GetAll())
	router.POST("/products/add", handler.Save())

	router.Run()
}

package main

import (
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/personas"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handler.NewProduct(service)
	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	router.Run()
}

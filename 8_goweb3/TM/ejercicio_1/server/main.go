package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/7_goweb2/Tarde/ejercicio_2/server/handler"

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

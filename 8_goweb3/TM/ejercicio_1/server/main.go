package main

import (
	products "github.com/extmatperez/meli_bootcamp2/tree/parra_diego/8_goweb3/TM/ejercicio_1/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/8_goweb3/TM/ejercicio_1/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	//
	router.PUT("/products/:id", controller.Update())

	router.Run()
}

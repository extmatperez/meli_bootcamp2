package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/internal/producto"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repository := producto.NewRepository()
	service := producto.NewService(repository)
	controller := handler.NewProduct(service)

	router.GET("/products/load", controller.LoadFile())
	router.GET("/products", controller.GetAll())
	router.POST("/products/add", controller.Store())

	router.Run()
}

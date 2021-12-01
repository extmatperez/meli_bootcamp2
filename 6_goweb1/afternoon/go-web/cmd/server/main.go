package main

import (
	handler "github.com/extmatperez/meli_bootcamp2/6_goweb1/afternoon/go-web/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/6_goweb1/afternoon/go-web/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := producto.NewRepository()
	service := producto.NewService(repo)
	controller := handler.NewProducto(service)

	router := gin.Default()

	router.GET("/productos/", controller.GetAll())
	router.POST("productos", controller.Store())

	router.Run(":8080")
}

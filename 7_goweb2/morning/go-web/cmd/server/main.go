package main

import (
	handler "github.com/extmatperez/meli_bootcamp2/7_goweb2/morning/go-web/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/7_goweb2/morning/go-web/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := productos.NewRepository()
	service := productos.NewService(repo)
	controller := handler.NewProducto(service)

	router := gin.Default()

	router.GET("/productos/", controller.GetAll())
	router.POST("productos", controller.Store())
	router.PUT("productos/:id", controller.Update())
	router.PATCH("productos/:id", controller.UpdateName())
	router.DELETE("productos/:id", controller.Delete())

	router.Run(":8080")
}

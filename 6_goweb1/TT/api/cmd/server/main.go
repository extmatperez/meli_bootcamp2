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

	routerGroup := router.Group("/products")
	{
		routerGroup.GET("/load", controller.LoadFile())
		routerGroup.GET("", controller.GetAll())
		routerGroup.POST("/add", controller.Store())
		routerGroup.DELETE("/:id", controller.Delete())
		routerGroup.PUT("/:id", controller.Update())
		routerGroup.PATCH("/:id", controller.UpdateNombre())
	}

	router.Run()
}

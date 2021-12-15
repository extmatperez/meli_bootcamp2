package main

import (
	handler "github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	store := store.NewStore(store.FileType, "/Users/fcenteno/Documents/GoFirst/meli_bootcamp2/11_testing2/proyecto/files/products.json")
	repo := productos.NewRepository(store)
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

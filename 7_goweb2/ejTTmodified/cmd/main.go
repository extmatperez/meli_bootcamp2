package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/7_goweb2/ejTTmodified/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/7_goweb2/ejTTmodified/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/7_goweb2/ejTTmodified/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	storage := store.NewStore("file", "products.json")
	repository := internal.NewRepository(storage)
	service := internal.NewService(repository)
	producto := handler.NewProducto(service)
	router.GET("/producto", producto.GetAll())
	router.POST("/producto", producto.Store())
	router.GET("/producto/:id", producto.GetProductById())
	router.PUT("/producto/:id", producto.Update())
	router.DELETE("/producto/:id", producto.Delete())
	router.PATCH("/producto/:id", producto.UpdateNombrePrecio())

	router.Run()

}

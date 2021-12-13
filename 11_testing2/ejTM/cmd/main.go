package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/11_testing2/ejTM/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/11_testing2/ejTM/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/pescie_juan/11_testing2/ejTM/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldnt load the environment")
	}
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

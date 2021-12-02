package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/products"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	db := store.New(store.FileType, "./productsList.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)
	router := gin.Default()

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	router.PUT("/products/:id", controller.Update())
	router.PATCH("/products/:id", controller.UpdateProd())
	router.DELETE("/products/:id", controller.Delete())
	router.Run()
}

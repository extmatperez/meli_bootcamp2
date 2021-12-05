package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/project/cmd/server/handler"
	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/project/internal/products"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/project/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	router.PUT("/products/:id", controller.Update())
	// router.PATCH("/products/:id", controller.UpdateProd())
	// router.DELETE("/products/:id", controller.Delete())
	router.Run()

}

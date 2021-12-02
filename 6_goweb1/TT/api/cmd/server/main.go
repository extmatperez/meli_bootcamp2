package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/internal/producto"
	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "../../internal/producto/products.json")

	repository := producto.NewRepository(db)
	service := producto.NewService(repository)
	controller := handler.NewProduct(service)

	routerGroup := router.Group("/products")
	{
		routerGroup.GET("", controller.GetAll())
		routerGroup.POST("/add", controller.Store())
		routerGroup.DELETE("/:id", controller.Delete())
		routerGroup.PUT("/:id", controller.Update())
		routerGroup.PATCH("/:id", controller.UpdateNombre())
	}

	router.Run()
}

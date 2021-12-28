package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/17_storage1/go_web/cmd/server/handler"
	internal "github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/17_storage1/go_web/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/17_storage1/go_web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./productos.json")

	repo := internal.NewRepository(db)
	service := internal.NewService(repo)
	controller := handler.NewProducto(service)

	trans := router.Group("/productos")
	{
		trans.GET("/get", controller.GetAll())
		trans.POST("/add", controller.Store())
		router.GET("/:id", controller.GetProductById())
		router.PUT("/:id", controller.Update())
		router.PATCH("/:id", controller.UpdateNombrePrecio())
		router.DELETE("/:id", controller.Delete())
	}

	router.Run()

}

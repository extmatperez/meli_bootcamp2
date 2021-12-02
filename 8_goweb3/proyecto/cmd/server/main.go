package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/cmd/server/handler"
	producto "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error: No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./productos.json")
	repo := producto.NewRepository(db)
	service := producto.NewService(repo)
	controller := handler.NewProducto(service)

	router.GET("/productos", controller.GetAll())
	router.POST("/addProductos", controller.Store())
	router.PUT("/modify/:id", controller.Modify())
	router.PATCH("/modifyNaPr/:id", controller.ModifyNamePrice())
	router.DELETE("/delete/:id", controller.Delete())

	router.Run()
}

package main

import (
	"log"

	products "github.com/extmatperez/meli_bootcamp2/tree/parra_diego/8_goweb3/TT/ejercicio_1/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/8_goweb3/TT/ejercicio_1/server/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}
	router := gin.Default()

	db := store.New(store.FileType, "./personasSalida.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	repo := products.NewRepository()
	service := products.NewService(repo)
	controller := handler.NewProduct(service)

	router.GET("/products/get", controller.GetAll())
	router.POST("/products/add", controller.Store())
	//
	router.PUT("/products/:id", controller.Update())
	router.DELETE("/products/:id", controller.Delete())
	router.PATCH("/products/:id", controller.UpdateNamePrice())

	router.Run()
}

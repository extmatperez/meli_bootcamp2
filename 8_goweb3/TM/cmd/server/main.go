package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/cmd/server/handler"
	payments "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/internal/payments"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/pkg/store"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./payments.json")
	repository := payments.NewRepository(db)
	service := payments.NewService(repository)
	controller := handler.NewPayment(service)

	payments := router.Group("/payments")
	{
		payments.GET("/get", controller.GetAll())
		payments.GET("/filter", controller.Filtrar())
		payments.POST("/", controller.Store())
		payments.PUT("/:id", controller.Update())
		payments.PATCH("/code/:id", controller.UpdateCodigo())
		payments.PATCH("/amount/:id", controller.UpdateMonto())
		payments.DELETE("/:id", controller.Delete())
	}

	router.Run(":8080")
}

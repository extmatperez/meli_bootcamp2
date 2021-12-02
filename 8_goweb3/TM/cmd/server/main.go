package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/cmd/server/handler"
	payments "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/internal/payments"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repository := payments.NewRepository()
	service := payments.NewService(repository)
	controller := handler.NewPayment(service)

	payments := router.Group("/payments")
	{
		payments.GET("/get", controller.GetAll())
		payments.GET("/filter", controller.Filter())
		payments.POST("/", controller.Store())
		payments.PUT("/:id", controller.Update())
		payments.PATCH("/code/:id", controller.UpdateCodigo())
		payments.PATCH("/amount/:id", controller.UpdateMonto())
		payments.DELETE("/:id", controller.Delete())
	}

	router.Run(":8080")
}

package main

import (
	"github.com/rossi_juancruz/meli_bootcamp2/7_goweb2/afternoon/proyecto/cmd/server/handler"
	personas "github.com/rossi_juancruz/meli_bootcamp2/7_goweb2/afternoon/proyecto/internal/personas"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := personas.NewRepository()
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	personasEP := r.Group("/personas")
	{
		personasEP.GET("/", controller.GetAll())
		personasEP.POST("/add", controller.Store())
		personasEP.PUT("/update/:id", controller.Update())
		personasEP.PATCH("/updateParcial/:id", controller.UpdateNombre())
		personasEP.DELETE("/delete/:id", controller.Delete())
	}

	r.Run()
}
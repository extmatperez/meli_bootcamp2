package main

import (
	handler "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/7_goweb2/TT/cmd/server/handler"
	personas "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/7_goweb2/TT/internal/personas"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := personas.NewRepository()
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())

	router.Run()
}

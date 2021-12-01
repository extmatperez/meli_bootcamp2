package main

import (
	"github.com/extmatperez/tree/zamora_damian/go-web/clase2/goModularizadoEnCapas/cmd/server/handler"
	personas "github.com/extmatperez/tree/zamora_damian/go-web/clase2/goModularizadoEnCapas/cmd/server/handler"
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

package main

import (
	//usuarios "github.com/extmatperez/meli_bootcamp2/tree/vargas_ignacio/7_goweb2/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	controller := usuarios.NewUsuario(service)

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())

	router.Run(":8000")
}

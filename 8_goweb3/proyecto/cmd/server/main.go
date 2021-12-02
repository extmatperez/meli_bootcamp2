package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/internal/transacciones"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	r := gin.Default()
	tr := r.Group("/transacciones")
	tr.POST("/add", t.Store())
	tr.GET("/get", t.GetAll())
	tr.GET("/load", t.Load())
	tr.PUT("/update/:id", t.Update())

	r.Run()
}

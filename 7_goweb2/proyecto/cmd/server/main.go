package main

import (
	transacciones "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transacciones"
	"github.com/gin-gonic/gin"
	handler "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/cmd/server/handler"
)

func main() {
	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	p:=handler.NewTransaccion(service)

	r:=gin.Default()
	tr := r.Group("/transacciones")
	tr.POST("/", p.Store())
	tr.GET(("/"), p.GetAll())

	r.Run()

}
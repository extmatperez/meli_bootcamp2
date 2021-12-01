package main

import (
	transacciones "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transacciones"
	"github.com/gin-gonic/gin"
	handler "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/cmd/server/handler"
)

func main() {
	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	t:=handler.NewTransaccion(service)

	r:=gin.Default()
	tr := r.Group("/transacciones")
	tr.POST("/", t.Store())
	tr.GET(("/"), t.GetAll())
	tr.PUT(":/id", t.Update())
	tr.PATCH(":/id", t.UpdateEmisor())
	tr.DELETE(":/id", t.Delete())

	r.Run()

}
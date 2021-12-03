package main

import (
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/internal/transacciones"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	r := gin.Default()
	tr := r.Group("/transacciones")
	tr.POST("/add", t.Store())
	tr.GET("/get", t.GetAll())
	tr.GET("/load", t.Load())
	tr.GET("/find/:id", t.FindById())
	tr.GET("/filter", t.FilterBy())
	tr.PUT("/update/:id", t.Update())
	tr.PATCH("/cod/:id", t.UpdateCod())
	tr.PATCH("/mon/:id", t.UpdateMon())
	tr.DELETE("/del/:id", t.Delete())

	r.Run()
}

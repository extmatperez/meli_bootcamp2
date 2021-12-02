package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/8_goweb3/PracticaTM/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/8_goweb3/PracticaTM/internal/transacciones"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	transac := handler.NewTransaccion(service)

	groupTransac := router.Group("/transacciones")

	//Endpoints GET
	groupTransac.GET("/", transac.GetAll())

	groupTransac.GET("/:id", transac.Search())
	groupTransac.GET("/filtros", transac.Filter())

	// //Endpoints POST
	groupTransac.POST("/cargar", transac.Store())

	// Endpoints Put
	groupTransac.PUT("/:id", transac.Update())

	// Endpoint delete
	groupTransac.DELETE("/:id", transac.Delete())

	// Endpoint Patch
	groupTransac.PATCH("/:id", transac.UpdateCodigoYMonto())

	router.Run()
}

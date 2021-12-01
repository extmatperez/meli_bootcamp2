package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/7_goweb2/PracticaTT/Ejercicio1/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/7_goweb2/PracticaTT/Ejercicio1/internal/transacciones"
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
	// groupTransac.GET("/filtros", filtrarTransacciones)

	// //Endpoints POST
	groupTransac.POST("/cargar", transac.Store())

	router.Run()
}

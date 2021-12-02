package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/9_goweb4/PracticaTM/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/9_goweb4/PracticaTM/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/9_goweb4/PracticaTM/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./transaccionesSalida.json")
	repo := transacciones.NewRepository(db)
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

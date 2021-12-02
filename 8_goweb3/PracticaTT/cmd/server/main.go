package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/8_goweb3/PracticaTT/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/8_goweb3/PracticaTT/internal/transacciones"
	database "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/8_goweb3/PracticaTT/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := database.New("string", "PruebaFile.json")
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

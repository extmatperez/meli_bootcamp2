/*
Configurar para que el token sea tomado de las variables de entorno al momento de realizar
la validación, para eso se deben realizar los siguientes pasos:
	1. Configurar la aplicación para que tome los valores que se encuentran en el archivo .env como variable de entorno.
	2. Quitar el valor del token del código y agregar como variable de entorno.
	3. Acceder al valor del token mediante la variable de entorno.
*/

package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/cmd/server/handler"
	productos "github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/TT/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "/Users/beconti/Desktop/meli_bootcamp2/7_goweb2/productos.json")

	repository := productos.NewRepository(db)
	service := productos.NewService(repository)
	controller := handler.NewProducto(service)

	productos := router.Group("/productos")
	{
		productos.GET("/", controller.GetAll())
		productos.POST("/", controller.Store())
		productos.PUT("/:id", controller.Update())
		productos.DELETE("/:id", controller.Delete())
		productos.PATCH("/:id", controller.UpdateNombrePrecio())
	}

	router.Run("localhost:8080")
}

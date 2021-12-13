package main

import (
	"log"

	handler "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/cmd/server/handler"
	transacciones "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html




func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error al intentar cargar el archivo .env")
	}
	db := store.New(store.FileType, "/transacciones.json")
	repo := transacciones.NewRepository(db)
	service := transacciones.NewService(repo)
	t := handler.NewTransaccion(service)

	r := gin.Default()
	tr := r.Group("/transacciones")
	tr.POST("/", t.Store())
	tr.GET(("/"), t.GetAll())
	tr.PUT(":/id", t.Update())
	tr.PATCH(":/id", t.UpdateEmisor())
	tr.DELETE(":/id", t.Delete())

	r.Run()

}

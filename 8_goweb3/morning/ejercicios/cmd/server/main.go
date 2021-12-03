package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/cmd/server/handler"
	personas "github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/personas"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	r := gin.Default()

	db := store.New(store.FileType, "./personasSalida.json")
	repo := personas.NewRepository(db)
	service := personas.NewService(repo)
	controller := handler.NewPersona(service)

	personasEP := r.Group("/personas")
	{
		personasEP.GET("/", controller.GetAll())
		personasEP.POST("/add", controller.Store())
		personasEP.PUT("/update/:id", controller.Update())
		personasEP.PATCH("/updateParcial/:id", controller.UpdateNombre())
		personasEP.DELETE("/delete/:id", controller.Delete())
	}

	r.Run()
}
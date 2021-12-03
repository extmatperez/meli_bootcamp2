package main

import (
	"log"

	handler "github.com/extmatperez/meli_bootcamp2/tree/Saavedra-Benjamin/8_goweb3/practica/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/tree/Saavedra-Benjamin/8_goweb3/practica/internal/users"
	"github.com/extmatperez/meli_bootcamp2/tree/Saavedra-Benjamin/8_goweb3/practica/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./personasSalida.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	router.GET("/personas/get", controller.GetAll())
	router.POST("/personas/add", controller.Store())
	router.PUT("/personas/:id", controller.Update())
	router.PATCH("/personas/:id", controller.UpdateName())
	router.DELETE("/personas/:id", controller.Delete())

	router.Run()
}

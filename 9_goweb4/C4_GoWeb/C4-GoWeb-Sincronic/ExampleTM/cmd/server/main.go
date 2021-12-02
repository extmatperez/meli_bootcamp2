package main

import (
	"log"

	"github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTM/cmd/server/handler"
	users "github.com/extmatperez/meli_bootcamp2/9_goweb4/C4_GoWeb/C4-GoWeb-Sincronic/ExampleTM/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Dont open file .env")
	}

	router := gin.Default()

	repo := users.NewRepository()
	service := users.NewService(repo)
	controller := handler.NewUser(service)

	router.GET("/users/get", controller.GetAll())
	router.POST("/users/add", controller.Store())
	router.PUT("/users/", controller.Update())
	router.Run()
}

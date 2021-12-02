package main

import (
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	// Routers
	router := gin.New()
	router.GET("/users")
	router.POST("/users/add")
	router.Run()
}
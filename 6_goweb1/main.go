package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int
	Name        string
	Color       string
	Price       float64
	Stock       int
	Code        int
	IsPublished bool
	CreatedAt   string
}

func main() {
	router := gin.Default()

	//handlers
	sayHello := func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Ro :D"})
	}

	getAll := func(c *gin.Context) {
		products := []Product{}
		data, err := os.ReadFile("./tematica.json") //recupero datos de tematica.json
		if err != nil {
			json.Unmarshal(data, &products)
		}
		c.JSON(200, products)
	}

	//ruta que saluda
	router.GET("hola", sayHello)

	//ruta que devuelve un listado de productos
	router.GET("products", getAll)

	router.Run()
}

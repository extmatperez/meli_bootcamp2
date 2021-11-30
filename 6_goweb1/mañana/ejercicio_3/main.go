package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id      int
	Name    string
	Color   string
	Price   float64
	Stock   int
	Code    string
	Publish bool
	Date    string
}

func GetAll(c *gin.Context) {
	archivo, _ := os.ReadFile("./products.json")
	var listProducts []Product
	json.Unmarshal(archivo, &listProducts)
	c.JSON(http.StatusOK, gin.H{
		"products": listProducts,
	})
}

func main() {

	s := gin.New()
	s.GET("/saludo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola Diego!!",
		})
	})

	s.GET("/productos", GetAll)
	s.Run()

}

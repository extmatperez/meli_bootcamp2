package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      int     `json:"code"`
	Published string  `json:"published"`
	Created   string  `json:"created"`
}

///// FUNCIONES HANDLERS ///////

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Estefi!",
	})
}

func getAll(c *gin.Context) {

	var prodList []Products
	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal([]byte(readProducts), &prodList); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"data": prodList,
	})
}

func main() {

	router := gin.Default()

	router.GET("/hello", sayHello)

	router.GET("/products", getAll)

	router.Run()
}

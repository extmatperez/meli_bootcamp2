package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

func readData() []Products {

	var list []Products
	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal([]byte(readProducts), &list); err != nil {
		log.Fatal(err)
	}
	return list
}

func getAll(c *gin.Context) {

	var prodList = readData()

	c.JSON(200, gin.H{
		"data": prodList,
	})
}
func main() {
	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Franco!",
		})
	})
	router.GET("/products", getAll)
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()

}

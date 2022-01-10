package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float32 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creation_date"`
}

func GetAllProducts(c *gin.Context) {
	prods, _ := os.ReadFile("./products.json")
	var products []Product
	unm := json.Unmarshal(prods, &products)
	if unm != nil {
		c.JSON(500, "Error.")
	} else {
		c.JSON(200, products)
	}
}

func main() {
	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola Ernesto",
		})
	})

	router.GET("/products", GetAllProducts)

	router.Run()
}

package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Color      string `json:"color"`
	Stock      int    `json:"stock"`
	Code       string `json:"code"`
	Published  bool   `json:"published"`
	Created_at string `json:"created_at"`
}

func getAllProducts(ctx *gin.Context) {
	bytes, err := os.ReadFile("../Ejercicio1/products.json")

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "No se pudieron cargar los productos",
		})
		return
	}

	var products []Product
	errUnmarshal := json.Unmarshal(bytes, &products)

	if errUnmarshal != nil {
		ctx.JSON(500, gin.H{
			"error": "Error parseando el JSON de productos",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"products": products,
	})
}

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hola " + name,
		})
	})

	router.GET("/products", getAllProducts)

	router.Run()
}

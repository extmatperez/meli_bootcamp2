package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Stock        int    `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creationDate"`
}

func main() {
	router := gin.Default()

	router.GET("/hola", func(ctx *gin.Context) {
		name := ctx.Query("name")
		message := "Hola "

		if message != "" {
			message = message + name
		}

		ctx.JSON(200, gin.H{
			"message": message,
		})

	})

	router.GET("/hola/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.JSON(200, gin.H{
			"message": "hola" + name,
		})
	})

	router.GET("/products", getProd)

	router.Run()

}
func getProd(ctx *gin.Context) {

	data, err := os.ReadFile("./products.json")

	var products []Product

	if err == nil {
		ctx.JSON(400, gin.H{
			"message": "No se encontro resultado",
		})
	}

	prod := json.Unmarshal(data, &products)
	ctx.JSON(200, prod)

	if prod != nil {
		ctx.JSON(400, gin.H{
			"message": "No se encontro resultado",
		})
	}
	ctx.JSON(200, gin.H{
		"products": products,
	})

}

package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Code        int     `json:"code" binding:"required"`
	IsPublished bool    `json:"ispublished" binding:"required"`
	CreatedAt   string  `json:"createdat" binding:"required"`
}

//func verifyToken(token int) (bool, string) {
//	return false, ""
//}

func main() {
	router := gin.Default()

	//handlers
	sayHello := func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Ro :D"})
	}

	getAll := func(c *gin.Context) {
		products := []Product{}
		data, err := os.ReadFile("./tematica.json") //recupero datos de tematica.json
		if err == nil {
			json.Unmarshal(data, &products)
		}
		c.JSON(200, products)
	}

	postOne := func(c *gin.Context) {
		// lo hice postear en el archivo porque lei mal el enunciado :P era en variable local
		//falta verificar bien los cammpos del product que llega por body
		token := c.GetHeader("token")
		if token == "secure" {
			products := []Product{}
			data, err := os.ReadFile("./tematica.json") //recupero datos de tematica.json
			if err == nil {
				json.Unmarshal(data, &products)

				var product Product
				err := c.ShouldBindJSON(&product)
				if err != nil {
					//tendria que recuperar cada campo
					c.JSON(400, gin.H{
						"error": "el campo x es requerido",
					})
				} else {
					if len(products) == 0 {
						product.ID = 1
					} else {
						product.ID = products[len(products)-1].ID + 1
					}
					products = append(products, product)
					data, _ = json.Marshal(&products)
					os.WriteFile("./tematica.json", data, 0644)
					c.JSON(200, product)
				}
			} else {
				c.JSON(404, gin.H{"error": "productos no encontrados"})
			}
		} else {
			c.JSON(401, gin.H{"error": "no tiene permisos para realizar la operacion solicitada"})
		}
	}

	//ruta que saluda
	router.GET("hola", sayHello)

	//ruta que devuelve un listado de productos
	router.GET("products", getAll)

	//ruta post
	router.POST("postProduct", postOne)

	router.Run()
}

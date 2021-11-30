package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
1. Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
2. Luego genera la lógica de filtrado de nuestro array.
3. Devolver por el endpoint el array filtrado.
*/

type Product struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fecha_de_creacion"`
}

func Ejemplo(context *gin.Context) {
	header := context.Request.Header
	fmt.Println("Header", header)
}
func Name(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
func GetID(c *gin.Context) {
	id := c.Param("id")
	var products []Product
	jsonFile, err := os.ReadFile("./products.json")
	if err == nil {
		json.Unmarshal(jsonFile, &products)
		idInt, _ := strconv.Atoi(id)
		flag := false
		for _, product := range products {
			if product.ID == idInt {
				flag = true
				c.JSON(http.StatusOK, product)
			}
		}
		if flag == false {
			c.String(404, "404 producto no encontrado")
		}
	} else {
		c.String(200, "404 producto no encontrado")
	}
}

func GetAll(c *gin.Context) {
	var products []interface{}
	jsonFile, err := os.ReadFile("./products.json")
	if err == nil {
		json.Unmarshal(jsonFile, &products)
		c.JSON(http.StatusOK, products)
		// c.JSON(http.StatusOK, gin.H{
		// 	"products": listProducts,
		// })
	}
}

func main() {
	router := gin.Default()
	router.GET("/products/:name", Name)
	router.GET("/producto/:id", GetID)
	router.GET("/products/All", GetAll)
	router.Run(":8080")
}

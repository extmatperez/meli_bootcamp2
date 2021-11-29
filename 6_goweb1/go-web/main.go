package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}

func saludar(c *gin.Context) {
	name := c.Param("fulano")
	c.JSON(200, gin.H{
		"message": "Hola " + name,
	})
}
func GetAll(c *gin.Context) {
	var products []Products
	ps, err := os.ReadFile("./6_goweb1/go-web/archivos/products.json")
	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(ps, &products)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": products,
	})
}

func main() {

	router := gin.Default()
	router.GET("/hola/:fulano", saludar)
	router.GET("/products", GetAll)
	router.Run()
}

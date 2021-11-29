package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Color     string `json:"color"`
	Precio    string `json:"precio"`
	Stock     int    `json:"stock"`
	Codigo    string `json:"codigo"`
	Publicado bool   `json:"publicado"`
	Creado    string `json:"creado"`
}

func main() {
	router := gin.Default()
	router.GET("/hola", saludar)
	router.GET("/productos", getAll)
	router.Run()
}

func saludar(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hola, Matias",
	})
}

func getAll(c *gin.Context) {
	productos, _ := os.ReadFile("prosductos.json")
	var punteroProductos []Producto
	err := json.Unmarshal(productos, &punteroProductos)
	if err != nil {
		c.JSON(500, "Ha ocurrido un error")
		return
	}
	c.JSON(200, punteroProductos)
}

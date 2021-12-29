package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/holaDiego", holaDiego)
	products := router.Group("/products")
	{
		products.GET("/productos/all", get_all)
		products.GET("/productos/filtered", productosFiltrados)
	}

	router.Run()

}

type Productos struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

func convert_json_to_productos() []Productos {
	jsonFile, err := os.ReadFile("usuarios.json")
	if err != nil {
		fmt.Println(err)
	}
	var products []Productos
	json.Unmarshal(jsonFile, &products)
	return products
}

func get_all(c *gin.Context) {
	products := convert_json_to_productos
	c.JSON(200, gin.H{
		"message": products,
	})
}

func holaDiego(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola Diego",
	})
}

func productosFiltrados(c *gin.Context) {
	products := convert_json_to_productos()
	fmt.Println(products)
}

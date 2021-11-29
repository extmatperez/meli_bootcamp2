package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}

func getAll(c *gin.Context) {
	data, err := os.ReadFile("../Ejercicio1/products.json")
	var ProductsRead []Product
	json.Unmarshal(data, &ProductsRead)

	if err != nil {
		fmt.Print("Error leyendo el archivo")
	}
	c.JSON(http.StatusOK, gin.H{
		"mensaje": ProductsRead,
	})
}

func main() {

	router := gin.Default()

	router.GET("/products", getAll)
	router.Run()
}
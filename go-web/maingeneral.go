//EJERCICIO 1 AL 3 C1-GO WEB - TM

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

type Product struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

func main() {
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		router := gin.Default()
		router.GET("/product", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": product,
			})
		})
		router.Run()
	}
	// ACEPTAR UN PARAMETRO POR URL
	/* router.GET("/productname/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hello %s", nombre)
	}) */
}

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID            string `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

func GetAll(c *gin.Context) {

	content, err := os.ReadFile("products.json")

	if err != nil {
		fmt.Println(err)
		panic("el archivo indicado no fue encontrado o está dañado\n")
	}

	p := []Product{}

	json.Unmarshal(content, &p)

	c.JSON(200, p)

}

func main() {
	router := gin.Default()

	router.GET("/hello/:name", func(c *gin.Context) {
		a1 := "Hola " + c.Param("name")
		//	query := c.Request.URL.Query()
		//	fmt.Println(query)
		c.JSON(200, gin.H{
			"message": a1,
		})
	})

	router.GET("/products", GetAll)

	router.Run()
}

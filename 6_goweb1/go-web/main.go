package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

func saludo(c *gin.Context) {
	nombre := c.Param("nombre")
	c.JSON(200, gin.H{
		"mensaje": "Hola, " + nombre,
	})
}
func GetAll(c *gin.Context) {

	var productosListo []Producto

	dbproductos, _ := ioutil.ReadFile("products.json")
	err := json.Unmarshal(dbproductos, &productosListo)

	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, productosListo)
	}
}
func main() {
	router := gin.Default()

	router.GET("/hola/:nombre", saludo)
	router.GET("/productos", GetAll)
	router.Run()
}

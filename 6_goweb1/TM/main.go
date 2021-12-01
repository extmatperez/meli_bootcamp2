package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {

	type productos struct {
		Id             int
		Nombre         string
		Color          string
		Precio         float64
		Stock          int
		Codigo         string
		Publicado      bool
		Fecha_creacion string
	}

	router := gin.Default()

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world Cristian",
		})
	})

	router.GET("/productos/GetAll", func(c *gin.Context) {
		datos, err := ioutil.ReadFile("../productos.json")

		if err == nil {
			datosString := string(datos)
			c.JSON(200, gin.H{
				"message": "se trajeron los datos",
				"datos":   datosString,
			})
		}

	})

	router.Run()
}

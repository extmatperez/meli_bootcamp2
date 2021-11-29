package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID            int
	Nombre        string
	Apellido      string
	Email         string
	Edad          int
	Altura        int
	Activo        bool
	FechaCreacion string
}

func main() {
	router := gin.Default()

	router.GET("/hello-world/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name)
		c.JSON(200, gin.H{
			"message": "Hola " + name + "!!",
		})
	})

	router.Run()

}

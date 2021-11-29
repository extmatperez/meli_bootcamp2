package main

import (
	"fmt"
	"os"

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

	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Nicolás!!",
		})
	})

	router.GET("/usuarios", func(c *gin.Context) {
		data, err := os.ReadFile("./users.json")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(data)
		}
		c.JSON(200, gin.H{
			"message": "Hola Nicolás!!",
		})
	})

	router.Run()

}

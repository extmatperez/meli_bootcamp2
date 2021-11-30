package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	person := person{"Dig", "Davila"}

	router.GET("/persons", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hola %v %v", person.Name, person.LastName),
		})
	})

	router.Run()
}

type person struct {
	Name     string
	LastName string
}

package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	per := persona{"Dig", "Davila"}

	var personas []persona
	personas = append(personas, per)
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"persona": per,
		})
	})

	router.Run()
}

type persona struct {
	Nombre   string
	Apellido string
}

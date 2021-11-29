package main

import "github.com/gin-gonic/gin"

type Persona struct {
	Name     string
	LastName string
}

func main() {
	router := gin.Default()

	per := Persona{"Jose", "Rios"}
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "hola")
	})
	router.GET("/hello2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"mensaje": per,
		})
	})

	router.Run()
}

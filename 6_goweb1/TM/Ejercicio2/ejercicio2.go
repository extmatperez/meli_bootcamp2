package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")

		message := "Hola"

		if name != "" {
			message = message + " " + name
		}

		c.JSON(200, gin.H{
			"message": message,
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hola " + name,
		})
	})

	router.Run()
}

package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	/* 	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola Ramiro",
		})
	}) */
	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "hello %s", name)
	})

	//inicializamos el router
	router.Run()
}

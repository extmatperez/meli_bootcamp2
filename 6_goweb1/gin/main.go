package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	router.Run() // Corremos nuestro server en el puerto, por defecto, el 8080. Si no, debemos darle dentro del metodo Run el puerto a usar.
	// router.Run(":8083") que pide que sea en el puerto 8083.
}

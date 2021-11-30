package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func saludo(c *gin.Context) {
	nombre := c.Param("nombre")
	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Hola " + nombre,
	})
}

func main() {
	s := gin.New()

	s.GET("/hola/:nombre", saludo)

	s.Run()
}

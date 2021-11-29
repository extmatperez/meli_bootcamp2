package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	nombre := "brian"
	s := gin.New()
	s.GET("/saludo", func(c *gin.Context) {
		c.String(http.StatusOK, "hola "+nombre)
	})
	s.Run()
}

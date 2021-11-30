package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id      int
	Name    string
	Color   string
	Price   float64
	Stock   int
	Code    string
	Publish bool
	Date    string
}

func saludo(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hola " + name,
	})

}

func main() {
	s := gin.New()
	s.GET("/saludo:name", saludo)
	s.Run()

}

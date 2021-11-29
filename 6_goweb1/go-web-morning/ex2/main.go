package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {
	router := gin.Default()
	
	p1 := &Person{"Juan Cruz", "Rossi"}

	router.GET("/hello-name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola " + p1.Name,
		})
	})

	router.Run()
}
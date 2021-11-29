package main

import (
	"github.com/gin-gonic/gin"
)

type Products struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      int     `json:"code"`
	Published string  `json:"published"`
	Created   string  `json:"created"`
}

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Estefi!",
		})
	})

	router.Run()
}

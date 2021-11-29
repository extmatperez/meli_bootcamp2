package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context){
		name,_ := c.GetQuery("name")
		c.JSON(http.StatusOK,gin.H{
			"message":"Hola " +name,
		})

	})

	router.Run()
 }



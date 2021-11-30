package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Age          int    `json:"age"`
	Height       int    `json:"height"`
	Active       bool   `json:"active"`
	CreationDate string `json:"creation_date"`
}

func main() {

	router := gin.Default()

	/* 	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola Ramiro",
		})
	}) */
	/* router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "hello %s", name)
	}) */
	router.GET("/users", func(c *gin.Context) {
		var userArr []Users
		readUsers, err := os.ReadFile("./users.json")

		if err != nil {
			fmt.Println(err)
		} else {
			json.Unmarshal(readUsers, &userArr)
		}
		c.JSON(200, gin.H{
			"users": userArr,
		})
		fmt.Println(userArr)
		fmt.Println(string(readUsers))
	})

	//inicializamos el router
	router.Run()
}

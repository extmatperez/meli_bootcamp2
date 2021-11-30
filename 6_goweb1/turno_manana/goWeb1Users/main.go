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
	CreationDate string `json:"date"`
}

func getAllProducts(c *gin.Context) {
	readUsers, err := os.ReadFile("./users.json")

	fmt.Println(string(readUsers))

	if err != nil {
		c.JSON(500, gin.H{
			"error": "No se pudieron cargar los usuarios",
		})
		return
	}

	var userArr []Users
	errUnmarshal := json.Unmarshal(readUsers, &userArr)

	if errUnmarshal != nil {
		c.JSON(500, gin.H{
			"error": "Error parseando el JSON de usuarios",
		})
		return
	}

	c.JSON(200, gin.H{
		"users": userArr,
	})
}

func main() {

	router := gin.Default()

	/* router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola Ramiro",
		})
	})
	router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "hello %s", name)
	}) */

	router.GET("/users", getAllProducts)

	//inicializamos el router
	router.Run()
}

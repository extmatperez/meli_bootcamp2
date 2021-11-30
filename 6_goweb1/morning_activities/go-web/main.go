package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creation_date"`
}

// Return all users
func GetAll(c *gin.Context) {
	var users_list []Users
	read_users, err := os.ReadFile("./users.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Users not found!",
		})
		//panic(err)
	} else {
		json.Unmarshal(read_users, &users_list)
		c.JSON(http.StatusOK, gin.H{
			"users": users_list,
		})
	}
	fmt.Println(users_list)
	fmt.Println(string(read_users))
}

func main() {
	router := gin.Default()
	// Return Hello World
	/* router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	}) */

	//Return name in params
	/* router.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name") // c.GetQuery
		c.String(http.StatusOK, "Hello %s!", name)
	}) */
	router.GET("/users", GetAll)
	router.Run()
}

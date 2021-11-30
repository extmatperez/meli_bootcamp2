package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int     `json: "id"`
	Name     string  `json: "name"`
	LastName string  `json: "last_name"`
	Email    string  `json: "email"`
	Age      int     `json: "age"`
	Height   float64 `json: "height"`
	Active   bool    `json: "active"`
	Created  string  `json: "created"`
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hi Federico"})
}

func getAll(c *gin.Context) {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var users []user
	json.Unmarshal(data, &users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func main() {

	s := gin.New()

	s.GET("/hi", hello)

	s.GET("/users", getAll)

	s.Run()
}

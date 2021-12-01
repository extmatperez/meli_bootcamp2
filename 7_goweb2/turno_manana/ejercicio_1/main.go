package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	Active    bool   `json:"active"`
	Date      string `json:"date"`
}

var user []Users

// Return all users
func get_users(c *gin.Context) {
	if len(user) == 0 {
		read_users, _ := os.ReadFile("./users.json")
		err := json.Unmarshal(read_users, &user)
		if err != nil {
			c.JSON(500, gin.H{
				"error": "Users not found!",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"users": user,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"users": user,
		})
	}
}
func post_users(c *gin.Context) {
	var user_id Users

	err := c.ShouldBindJSON(&user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		last_id := user[len(user)-1].ID + 1
		user_id.ID = last_id
		user = append(user, user_id)
		c.JSON(http.StatusOK, user_id)
	}
}

func main() {
	router := gin.Default()
	router.GET("/users", get_users)
	router.POST("/users", post_users)

	router.Run()
}

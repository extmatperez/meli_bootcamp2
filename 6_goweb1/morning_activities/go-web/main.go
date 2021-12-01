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

/* func readData() []Users {

	var list []Users
	read_users, _ := os.ReadFile("./Users.json")

	if err := json.Unmarshal([]byte(read_users), &list); err != nil {
		log.Fatal(err)
	}
	return list
} */

// Return Hello World!
func hello_world(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func hello_you(c *gin.Context) {
	name := c.Param("name") // c.GetQuery
	c.String(http.StatusOK, "Hello %s!", name)
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
}

/* func getbyName(c *gin.Context) {

	var user_list = GetAll()
	var filtered []Users

	for _, us := range user_list {
		if c.Query("name") == us.Name {
			filtered = append(filtered, us)
		}
	}

	c.JSON(200, gin.H{
		"response": filtered,
	})

} */

func main() {
	router := gin.Default()
	// Return Hello World
	router.GET("/", hello_world)

	//Return name in params
	router.GET("/users/:name", hello_you)

	// Return all users
	router.GET("/users", GetAll)

	/* usersfiltered := router.Group("/usersfiltered")
	{
		usersfiltered.GET("/name", getbyName)
	} */

	router.Run()
}

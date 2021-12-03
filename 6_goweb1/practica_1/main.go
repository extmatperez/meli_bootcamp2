package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	ID           int     `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Active bool    `json:"active"`
	CreationDate string  `json:"creation_date"`
}
func main() {
	// Routers
	router := gin.New()
	router.GET("/sayHello/:name", sayHello)
	router.GET("/users", GetALL)
	router.GET("/users/filter", GetFilter)
	router.Run()
}
func GetFilter(c *gin.Context) {
	//variables
	var (
		//users []Users
		id, _ = c.GetQuery("ID")
		//first_name, _ = c.GetQuery("firstName")
	)
	idInt, _ := strconv.Atoi(id)
	c.JSON(200,idInt)
}
func GetALL(c *gin.Context) {
	// Variables
	var (
		users []Users
	)
	// UnMarshal
	jsonFile, _ := os.Open("users.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &users)
	c.JSON(200, users)
}
func sayHello(c *gin.Context)  {
	// set param variable
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello, "+name,
	})
}
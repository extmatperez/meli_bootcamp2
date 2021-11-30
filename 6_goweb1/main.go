package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

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

func filterByString(field string, filter string) bool {
	if filter != "" {
		if !strings.EqualFold(field, filter) {
			return false
		}
	}
	return true
}

func filterByInt(field int, filter int) bool {
	if filter != 0 && field != filter {
		return false
	}
	return true
}

func filterByFloat(field float64, filter float64) bool {
	if filter != 0.0 && field != filter {
		return false
	}
	return true
}

func filterByBool(field bool, filter string) bool {
	if filter != "" {
		boolFilter, _ := strconv.ParseBool(filter)
		return field == boolFilter
	}

	return true
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hi Federico"})
}

func getAll(ctx *gin.Context) {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var filteredUsers, users []user
	err = json.Unmarshal(data, &users)

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Can't obtain users right now.",
		})
		return
	}

	nameQuery := ctx.Query("name")
	lastNameQuery := ctx.Query("last_name")
	emailQuery := ctx.Query("email")
	ageQuery, _ := strconv.Atoi(ctx.Query("age"))
	heightQuery, _ := strconv.ParseFloat(ctx.Query("height"), 64)
	activeQuery := ctx.Query("active")
	createdQuery := ctx.Query("created")

	for _, user := range users {
		addUser := true

		if !filterByString(user.Name, nameQuery) || !filterByString(user.LastName, lastNameQuery) {
			addUser = false
		}

		if !filterByString(user.Email, emailQuery) || !filterByString(user.Created, createdQuery) {
			addUser = false
		}

		if !filterByInt(user.Age, ageQuery) {
			addUser = false
		}

		if !filterByBool(user.Active, activeQuery) {
			addUser = false
		}

		if !filterByFloat(user.Height, heightQuery) {
			addUser = false
		}

		if addUser {
			filteredUsers = append(filteredUsers, user)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"users": filteredUsers})
}

func main() {

	s := gin.New()

	s.GET("/hi", hello)

	s.GET("/users", getAll)

	s.Run()
}

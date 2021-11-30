package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	LastName string  `json:"last_name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Age      int     `json:"age" binding:"required"`
	Height   float64 `json:"height" binding:"required"`
	Active   bool    `json:"active" binding:"required"`
	Created  string  `json:"created" binding:"required"`
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

func getUsers() ([]user, error) {
	data, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var users []user
	err = json.Unmarshal(data, &users)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func saveFile(users []user) error {
	text, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile("./users.json", text, 0644)
	if err != nil {
		return err
	}

	return nil
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hi Federico"})
}

func getAll(ctx *gin.Context) {
	users, err := getUsers()
	var filteredUsers []user

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

func getByID(ctx *gin.Context) {
	users, err := getUsers()

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Can't obtain users right now.",
		})
		return
	}

	idQuery, _ := strconv.Atoi(ctx.Param("id"))

	i := 0
	for i < len(users) && idQuery != users[i].ID {
		i++
	}

	if i < len(users) {
		ctx.JSON(http.StatusOK, gin.H{"user": users[i]})
	} else {
		ctx.JSON(http.StatusNotFound, "")
	}

}

func validateRequiredField(fieldName, err string) (string, bool) {
	if !strings.Contains(err, "'required'") {
		return "", true
	}

	if strings.Contains(err, "'"+fieldName+"'") {
		return fieldName, false
	}

	return "", true
}

func validateBodyUser(err error) string {
	var requiredFields []string
	requiredFields = append(requiredFields, "name", "lastName", "email", "age", "height", "active", "created")

	msg := ""

	for _, fieldName := range requiredFields {
		field, validated := validateRequiredField(fieldName, strings.ToLower(err.Error()))

		if !validated {
			fmt.Println(field)
			msg = fmt.Sprintf("%s %s", msg, field)
		}
	}

	return msg
}

func validateToken(ctx *gin.Context) error {
	token := ctx.GetHeader("token")

	if token != "123456" {
		return errors.New("Invalid Token")
	}

	return nil
}

func save(ctx *gin.Context) {

	errToken := validateToken(ctx)

	if errToken != nil {
		ctx.JSON(401, gin.H{
			"error": fmt.Sprintf("%s", errToken),
		})
		return
	}

	var newUser user
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		msg := validateBodyUser(err)

		if msg == "" {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
		return
	}

	users, errGet := getUsers()
	if errGet != nil {
		ctx.JSON(500, gin.H{
			"error": "Can't obtain users right now.",
		})
		return
	}

	id := users[len(users)-1].ID

	newUser.ID = id + 1
	users = append(users, newUser)

	errSave := saveFile(users)

	if errSave != nil {
		ctx.JSON(500, gin.H{
			"error": "Can't save the new user right now.",
		})
		return
	}

	ctx.JSON(200, newUser)
}

func main() {

	s := gin.New()

	s.GET("/hi", hello)

	group := s.Group("/users")
	{
		group.GET("", getAll)
		group.GET("/:id", getByID)
		group.POST("", save)
	}

	s.Run()
}

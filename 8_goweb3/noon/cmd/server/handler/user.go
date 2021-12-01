package handler

import (
	"fmt"
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/8_goweb3/noon/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`
	Email    string  `json:"email"`
	Age      int     `json:"age"`
	Height   float64 `json:"height"`
	Active   bool    `json:"active"`
	Created  string  `json:"created"`
}

type User struct {
	service users.Service
}

func NewUser(ser users.Service) *User {
	return &User{service: ser}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "Invalid token",
			})
			return
		}

		users, err := u.service.GetAll()

		if err != nil {
			ctx.JSON(400, "There was an error "+err.Error())
		} else {
			ctx.JSON(200, users)
		}
	}
}

func (u *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "Invalid token",
			})
			return
		}

		var user request
		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			ctx.JSON(400, gin.H{
				"Error": "There was an error when storing the user: " + err.Error(),
			})
			return
		}

		newUser, errStore := u.service.Store(user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.Created)
		if errStore != nil {
			ctx.JSON(404, gin.H{
				"Error": "There was an error when storing the user: " + errStore.Error(),
			})
			return
		}

		ctx.JSON(200, newUser)
	}
}

func (u *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "Invalid token",
			})
			return
		}

		id, errId := strconv.Atoi(ctx.Param("id"))
		if errId != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid ID",
			})
		}

		var user request
		err := ctx.ShouldBindJSON(&user)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "There was an error when storing the user: " + err.Error(),
			})
			return
		}

		checkMsg := validateUpdateFields(user)
		if checkMsg != "" {
			ctx.JSON(400, gin.H{
				"error": fmt.Sprintf("Required field/s missing: %s", checkMsg),
			})

			return
		}

		updatedUser, errStore := u.service.Update(id, user.Name, user.LastName, user.Email, user.Age, user.Height, user.Active, user.Created)
		if errStore != nil {
			ctx.JSON(404, gin.H{
				"Error": "There was an error when storing the user: " + errStore.Error(),
			})
			return
		}

		ctx.JSON(200, updatedUser)
	}
}

func validateUpdateFields(u request) string {
	msg := ""

	if u.Name == "" {
		msg = fmt.Sprintf("%s %s", "name")
	}

	if u.LastName == "" {
		msg = fmt.Sprintf("%s %s", "last_name")
	}

	if u.Email == "" {
		msg = fmt.Sprintf("%s %s", "email")
	}

	if u.Age == 0 {
		msg = fmt.Sprintf("%s %s", "age")
	}

	if u.Height == 0.0 {
		msg = fmt.Sprintf("%s %s", "height")
	}

	if u.Created == "" {
		msg = fmt.Sprintf("%s %s", "created")
	}

	return msg
}

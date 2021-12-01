package handler

import (
	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/7_goweb2/afternoon/internal/users"
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

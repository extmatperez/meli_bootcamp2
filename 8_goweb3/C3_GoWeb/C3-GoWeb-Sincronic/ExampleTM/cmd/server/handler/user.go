package handler

import (
	users "github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-Sincronic/ExampleTM/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type User struct {
	service users.Service
}

func NewUser(ser users.Service) *User {
	return &User{
		service: ser}
}

func (us *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		users, err := us.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error: %v", err)
		} else {
			ctx.JSON(200, users)
		}
	}
}

func (controller *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user request

		err := ctx.ShouldBind(&user)
		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona: %v", err)
		} else {
			controller.service.Store(user.FirstName, user.LastName, user.Age)
		}
	}
}

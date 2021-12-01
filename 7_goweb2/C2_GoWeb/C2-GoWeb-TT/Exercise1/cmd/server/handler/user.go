package handler

import (
	users "github.com/extmatperez/meli_bootcamp2/7_goweb2/C2_GoWeb/C2-GoWeb-TT/Exercise1/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Age         int    `json:"age"`
	Height      int    `json:"height"`
	Active      bool   `json:"active"`
	CrationDate string `json:"cration_date"`
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
			response, err := controller.service.Store(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
				// controller.service.Store(user.FirstName, user.LastName, user.Email, user.Age, user.Height, user.Active, user.CrationDate)
			}
		}
	}
}

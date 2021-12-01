package handler

import (
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/8_goweb3/C3_GoWeb/C3-GoWeb-TM/Exercise1/internal/users"
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

func (controller *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var us request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&us)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			usuarioUpdate, err := controller.service.Update(int(id), us.FirstName, us.LastName, us.Email, us.Age, us.Height, us.Active, us.CrationDate)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, usuarioUpdate)
			}
		}

	}
}

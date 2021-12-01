package handler

import (
	users "github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/7_goweb2/tarde/ejercicio1/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        string `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

type User struct {
	service users.Service
}

func NewUser(ser users.Service) *User {
	return &User{service: ser}
}

func (per *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		personas, err := per.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, personas)
		}
	}
}

func (controller *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var usr request

		err := ctx.ShouldBindJSON(&usr)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Store(usr.Nombre, usr.Apellido, usr.Email, usr.Edad, usr.Altura, usr.Activo, usr.FechaCreacion)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

package handler

import (
	"strconv"

	users "github.com/extmatperez/meli_bootcamp2/tree/brian_beltran/8_goweb3/tarde/ejercicio1/internal/users"
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

func (controller *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var usr request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&usr)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			userActualizada, err := controller.service.Update(int(id), usr.Nombre, usr.Apellido, usr.Email, usr.Edad, usr.Altura, usr.Activo, usr.FechaCreacion)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, userActualizada)
			}
		}

	}
}

func (controller *User) UpdateNombre() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var usr request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&usr)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			if usr.Nombre == "" {
				ctx.String(404, "El nombre no puede estar vacío")
				return
			}
			personaActualizada, err := controller.service.UpdateNombre(int(id), usr.Nombre)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, personaActualizada)
			}
		}

	}
}

func (controller *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, err.Error())
		} else {
			ctx.String(200, "La persona %d ha sido eliminada", id)
		}

	}
}

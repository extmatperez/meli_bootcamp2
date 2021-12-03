package handler

import (
	"os"
	"strconv"
	"time"

	users "github.com/extmatperez/meli_bootcamp2/tree/Saavedra-Benjamin/8_goweb3/practica/internal/users"
	"github.com/gin-gonic/gin"
)

type request struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Age       int     `json:"age"`
	Height    float64 `json:"height"`
}

type User struct {
	service users.Service
}

func NewUser(ser users.Service) *User {
	return &User{service: ser}
}

func validateToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.String(400, "Falta token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.String(404, "Token incorrecto")
		return false
	}

	return true
}

func (per *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validateToken(ctx) {
			return
		}

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

		if !validateToken(ctx) {
			return
		}

		var usr request

		err := ctx.ShouldBindJSON(&usr)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Store(usr.FirstName, usr.LastName, usr.Email, usr.Age, usr.Height)
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

		if !validateToken(ctx) {
			return
		}

		var usr request
		var (
			active       = true
			creationDate = time.Now()
		)

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&usr)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			personaActualizada, err := controller.service.Update(int(id), usr.FirstName, usr.LastName, usr.Email, usr.Age, usr.Height, active, creationDate)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, personaActualizada)
			}
		}

	}
}

func (controller *User) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validateToken(ctx) {
			return
		}

		var usr request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&usr)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			if usr.FirstName == "" {
				ctx.String(404, "El nombre no puede estar vacío")
				return
			}
			personaActualizada, err := controller.service.UpdateName(int(id), usr.FirstName)
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

		if !validateToken(ctx) {
			return
		}

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

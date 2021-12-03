package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	personas "github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/internal/personas"
	"github.com/rossi_juancruz/meli_bootcamp2/8_goweb3/morning/ejercicios/pkg/web"
)

type request struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

type Persona struct {
	service personas.Service
}

//Nueva persona
func NewPersona(ser personas.Service) *Persona {
	return &Persona{service: ser}
}

//Validar token de auth
func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "Falta token"))
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.JSON(400, web.NewResponse(400, nil, "Token invalido"))
		return false
	}

	return true
}

func (per *Persona) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !(validarToken(ctx)) {
			return
		}

		personas, err := per.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
		} else {
			ctx.JSON(200, web.NewResponse(200, personas, ""))
		}
	}
}

func (controller *Persona) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !(validarToken(ctx)) {
			return
		}

		var perso request

		err := ctx.ShouldBindJSON(&perso)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Store(perso.Nombre, perso.Apellido, perso.Edad)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (controller *Persona) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if !(validarToken(ctx)) {
			return
		}

		var per request
		err := ctx.ShouldBindJSON(&per)
		
		id, err_int := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err_int != nil {
			ctx.String(404, "Invalid ID")
		}
		
		if err != nil {
			ctx.String(404, "No se pudo actualizar la persona %v", err)
		} else {
			perActualizada, err := controller.service.Update(int(id), per.Nombre, per.Apellido, per.Edad)
			if err != nil {
				ctx.JSON(404, err.Error())
			} else {
				ctx.JSON(200, perActualizada)
			}
		}
	}
}

func (controller *Persona) UpdateNombre() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if !(validarToken(ctx)) {
			return
		}
	
		id, err_int := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err_int != nil {
			ctx.String(404, "Invalid ID")
		}
		
		var per request
		err := ctx.ShouldBindJSON(&per)

		if per.Nombre == "" {
			ctx.String(404, "el nombre no puede estar vacio")
			return
		}

		if err != nil {
			ctx.String(404, "No se pudo actualizar la persona %v", err)
		} else {
			_, err = controller.service.UpdateNombre(int(id), per.Nombre)
			if err != nil {
				ctx.JSON(404, err.Error())
			} else {
				ctx.String(200, "La persona %d ha sido actualizada", id)
			}
		}
	}
}

func (controller *Persona) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if !(validarToken(ctx)) {
			return
		}
	
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(404, "Invalid ID")
		}
		
		
			err = controller.service.Delete(int(id))
			if err != nil {
				ctx.JSON(404, err.Error())
			} else {
				msg := fmt.Sprintf("Deleted id %d successfully", id)
				ctx.JSON(200, gin.H{
					"message": msg,
				})
			}
		
	}
}
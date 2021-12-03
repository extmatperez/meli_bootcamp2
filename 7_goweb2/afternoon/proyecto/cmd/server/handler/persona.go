package handler

import (
	personas "github.com/rossi_juancruz/meli_bootcamp2/7_goweb2/afternoon/proyecto/internal/personas"
	"github.com/gin-gonic/gin"
	"strconv"
)

type request struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

type Persona struct {
	service personas.Service
}

func NewPersona(ser personas.Service) *Persona {
	return &Persona{service: ser}
}

func (per *Persona) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		personas, err := per.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, personas)
		}
	}
}

func (controller *Persona) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

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
				ctx.String(200, "La persona %d ha sido eliminada", id)
			}
		}
	}
}

func (controller *Persona) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {
	
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(404, "Invalid ID")
		}
		
		
			err = controller.service.Delete(int(id))
			if err != nil {
				ctx.JSON(404, err.Error())
			} else {
				ctx.JSON(200, err)
			}
		
	}
}
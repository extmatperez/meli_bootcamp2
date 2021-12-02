// ETAPA 3: este es el controller, va a tener todos los datos que yo recibo.
// Internamente voy a necesitar un servicio, y se lo tengo que asignar a un servicio.

// Para que el controller sea usado, en este caso, los dos metodos que tiene, deben devolver un gin.HandlerFunc.
// Nosotros para devolver eso

package handler

import (
	personas "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/7_goweb2/TT/internal/personas"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
}

type Persona struct {
	service personas.Service
}

func NewPersona(service personas.Service) *Persona {
	return &Persona{service: service}
}

func (per *Persona) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// RECORDAR EN ESTOS METODOS LLAMAR AL TOKEN. ESTAN EN LA PRESENTACION. token := ctx.Request.Header.Get("token")
		personas, err := per.service.GetAll()
		if err != nil {
			ctx.String(400, "Hubo un error: %v", err)
		} else {
			ctx.JSON(200, personas)
		}
	}
}

func (controller *Persona) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var personita request

		err := ctx.ShouldBindJSON(&personita)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona: %v", err)
		} else {
			response, err := controller.service.Store(personita.Nombre, personita.Apellido, personita.Edad)
			if err != nil {
				ctx.String(400, "No se puede cargar la persona: %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

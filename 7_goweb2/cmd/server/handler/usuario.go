package handler

import (
	//usuarios "github.com/extmatperez/meli_bootcamp2/7_goweb2/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(ser usuarios.Service) *Usuario {
	return &Usuario{service: ser}
}

func (usr *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		usuarios, err := usr.service.GetAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, usuarios)
		}
	}
}

func (controller *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var usu request

		err := ctx.ShouldBindJSON(&usu)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar una persona %v", err)
		} else {
			response, err := controller.service.Store(usu.Nombre, usu.Apellido, usu.Email, usu.Edad, usu.Altura, usu.Activo, usu.FechaCreacion)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

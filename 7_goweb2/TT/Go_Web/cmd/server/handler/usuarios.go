package handler

import (
	usuarios "github.com/extmatperez/meli_bootcamp2/tree/aponte_nicolas/7_goweb2/TT/Go_Web/internal/usuarios"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUsuario(serv usuarios.Service) *Usuario {
	return &Usuario{service: serv}
}

func (control *Usuario) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarios, err := control.service.GetAll()
		if err != nil {
			c.String(400, "Hubo un error %v", err)
		} else {
			c.JSON(200, usuarios)
		}
	}
}

func (control *Usuario) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser request
		err := c.ShouldBindJSON(&newUser)
		if err != nil {
			c.String(400, "Hubo un error al querer cargar un usuario %v", err)
		} else {
			response, err := control.service.Store(newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura, newUser.Activo, newUser.FechaCreacion)
			if err != nil {
				c.String(400, "No se pudo cargar la persona %v", err)
			} else {
				c.JSON(200, response)
			}
		}
	}
}

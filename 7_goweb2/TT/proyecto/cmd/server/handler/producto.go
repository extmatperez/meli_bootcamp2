package handler

import (
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/personas"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fecha_de_creacion"`
}

type Product struct {
	service products.Service
}

func NewProduct(prod products.Service) *Product {
	return &Product{
		service: prod,
	}
}
func (prod *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}
		p, err := prod.service.GetAll()
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)
	}
}
func (controller *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var prod request
		err := c.ShouldBindJSON(&prod)
		if err != nil {
			c.String(400, "Hubo un error al querer cargar la personas %v", err)
		} else {
			response, err := controller.service.Store(prod.Nombre, prod.Color, prod.Precio, prod.Stock, prod.Codigo, prod.Publicado, prod.FechaDeCreacion)
			if err != nil {
				c.String(400, "No se pudo cargar la persona %v", err)
			} else {
				c.JSON(200, response)
			}
		}
	}
}

package handler

import (
	producto
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre            string  `json:"nombre"`
	Color             string  `json:"color"`
	Precio            float64 `json:"precio"`
	Stock             int     `json:"stock"`
	Codigo            string  `json:"codigo"`
	Publicado         bool    `json:"publicado"`
	Fecha_de_creacion string  `json:"fecha_de_creacion"`
}

type Producto struct {
	service producto.Service
}

func NewProducto(ser producto.Service) *Producto {
	return &Producto{service: ser}
}

func (prod *Producto) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		prod, err := prod.service.GetAll()

		if err != nil {
			c.String(400, "Hubo un error %v", err)
		} else {
			c.JSON(200, prod)
		}
	}
}

func (controller *Producto) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var prod request
		err := c.ShouldBindJSON(&prod)
		if err != nil {
			c.String(400, "Hubo un error al querer cargar un producto %v", err)
		} else {
			response, err := controller.service.Store(prod.Nombre, prod.Codigo, prod.Color, prod.Fecha_de_creacion, prod.Precio, prod.Publicado, prod.Stock)

			if err != nil {
				c.String(400, "No se pudo cargar el producto: %v", err)
			} else {
				c.JSON(200, response)
			}
		}
	}
}
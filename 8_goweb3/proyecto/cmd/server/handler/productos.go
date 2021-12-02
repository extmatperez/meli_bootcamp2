package handler

import (
	"os"
	"strconv"

	producto "github.com/extmatperez/meli_bootcamp2/8_goweb3/proyecto/internal/productos"
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

func validarToken(c *gin.Context) bool {
	token := c.GetHeader("token")

	if token == "" {
		c.String(400, "Falta token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if tokenENV != token {
		c.String(400, "Token incorrecto")
		return false
	}
	return true
}

func (prod *Producto) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}

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

		if !validarToken(c) {
			return
		}
		var prod request
		err := c.ShouldBindJSON(&prod)
		if err != nil {
			c.String(400, "Hubo un error al querer cargar un producto %v", err)
		} else {
			response, err := controller.service.Store(prod.Stock, prod.Nombre, prod.Codigo, prod.Color, prod.Fecha_de_creacion, prod.Precio, prod.Publicado)

			if err != nil {
				c.String(400, "No se pudo cargar el producto: %v", err)
			} else {
				c.JSON(200, response)
			}
		}
	}
}

func (controller *Producto) Modify() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}
		var prod request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.String(404, "Hubo un error, el id es invalido")
		}

		err = c.ShouldBindJSON(&prod)

		if err != nil {
			c.String(400, "Hubo un error en el body")
		} else {
			prodActualizado, err := controller.service.Modify(int(id), prod.Stock, prod.Nombre, prod.Codigo, prod.Color, prod.Fecha_de_creacion, prod.Precio, prod.Publicado)
			if err != nil {
				c.JSON(400, err.Error())
			} else {
				c.JSON(200, prodActualizado)
			}
		}
	}
}

func (controller *Producto) ModifyNamePrice() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}
		var prod request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.String(404, "Hubo un error, el id es invalido")
		}

		err = c.ShouldBindJSON(&prod)

		if err != nil {
			c.String(400, "Hubo un error en el body")
		} else {

			prodActualizado, err := controller.service.ModifyNamePrice(int(id), prod.Nombre, prod.Precio)
			if err != nil {
				c.JSON(400, err.Error())
			} else {
				c.JSON(200, prodActualizado)
			}
		}
	}
}

func (controller *Producto) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !validarToken(c) {
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.String(404, "Hubo un error, el id es invalido")
		}

		err = controller.service.Delete(int(id))
		if err != nil {
			c.JSON(400, err.Error())
		} else {
			c.String(200, "El producto ha sido eliminado")
		}
	}
}

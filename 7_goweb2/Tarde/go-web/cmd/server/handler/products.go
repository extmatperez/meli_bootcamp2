package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/Tarde/go-web/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Nombre        string `json:"nombre" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Precio        int    `json:"precio" binding:"required"`
	Stock         int    `json:"stock" binding:"required"`
	Codigo        string `json:"codigo" binding:"required"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion" binding:"required"`
}

type requestPatchNamePrice struct {
	Nombre string `json:"nombre" binding:"required"`
	Precio int    `json:"precio" binding:"required"`
}

type Product struct {
	serv products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{serv: s}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "1234" {
			c.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
			return
		}

		prods, err := p.serv.GetAll()

		if err != nil {
			c.String(404, err.Error())
			return
		}
		c.JSON(200, prods)
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "1234" {
			c.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
			return
		}

		var newRequest request
		err := c.ShouldBindJSON(&newRequest)
		if err != nil {

			if strings.Contains(err.Error(), "required") {
				tipos := reflect.TypeOf(newRequest)
				i := 0
				var errores []string
				for i = 0; i < tipos.NumField(); i++ {
					if strings.Contains(err.Error(), tipos.Field(i).Name) {
						errores = append(errores, fmt.Sprintf("Error: el campo %s es requerido", tipos.Field(i).Name))
					}
				}
				if len(errores) == 1 {
					c.JSON(400, gin.H{
						"error": errores[0],
					})
				} else {
					c.JSON(400, errores)
				}
			} else {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			}

		} else {

			prod, err := p.serv.Store(newRequest.Nombre, newRequest.Color, newRequest.Precio, newRequest.Stock, newRequest.Codigo, newRequest.Publicado, newRequest.FechaCreacion)

			if err != nil {
				c.String(500, err.Error())
			} else {
				c.JSON(200, prod)
			}

		}

	}
}

func (p *Product) FindById() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "1234" {
			c.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "Id invalido")
			return
		}

		prod, err := p.serv.FindById(id)

		if err != nil {
			c.String(404, err.Error())
		} else {
			c.JSON(200, prod)
		}

	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "1234" {
			c.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
			return
		}
		var updateRequest request
		err := c.ShouldBindJSON(&updateRequest)
		if err != nil {
			if strings.Contains(err.Error(), "required") {
				tipos := reflect.TypeOf(updateRequest)
				i := 0
				var errores []string
				for i = 0; i < tipos.NumField(); i++ {
					if strings.Contains(err.Error(), tipos.Field(i).Name) {
						errores = append(errores, fmt.Sprintf("Error: el campo %s es requerido", tipos.Field(i).Name))
					}
				}
				if len(errores) == 1 {
					c.JSON(400, gin.H{
						"error": errores[0],
					})
				} else {
					c.JSON(400, errores)
				}
			} else {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "Id invalido")
			return
		}

		prod, err := p.serv.Update(id, updateRequest.Nombre, updateRequest.Color, updateRequest.Precio, updateRequest.Stock, updateRequest.Codigo, updateRequest.Publicado, updateRequest.FechaCreacion)

		if err != nil {
			c.String(404, err.Error())
			return
		}

		c.JSON(200, prod)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "1234" {
			c.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "Id invalido")
			return
		}

		err = p.serv.Delete(id)
		if err != nil {
			c.String(404, err.Error())
			return
		}

		c.String(200, "Eliminado exitosamente")

	}
}

func (p *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "1234" {
			c.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "Id invalido")
			return
		}

		var updateRequest requestPatchNamePrice
		err = c.ShouldBindJSON(&updateRequest)
		if err != nil {
			if strings.Contains(err.Error(), "required") {
				tipos := reflect.TypeOf(updateRequest)
				i := 0
				var errores []string
				for i = 0; i < tipos.NumField(); i++ {
					if strings.Contains(err.Error(), tipos.Field(i).Name) {
						errores = append(errores, fmt.Sprintf("Error: el campo %s es requerido", tipos.Field(i).Name))
					}
				}
				if len(errores) == 1 {
					c.JSON(400, gin.H{
						"error": errores[0],
					})
				} else {
					c.JSON(400, errores)
				}
			} else {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
			}
			return
		}

		prod, err := p.serv.UpdateNameAndPrice(id, updateRequest.Nombre, updateRequest.Precio)
		if err != nil {
			c.String(404, err.Error())
			return
		}

		c.JSON(200, prod)

	}
}

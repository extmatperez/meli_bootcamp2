package handler

import (
	"fmt"
	"os"
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/internal/products"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      int     `json:"code"`
	Published string  `json:"published"`
	Created   string  `json:"created"`
}

type Product struct {
	service products.Service
}

func NewProduct(serv products.Service) *Product {
	return &Product{service: serv}

}

func ValidateToken(c *gin.Context) bool {
	token := c.GetHeader("token")

	if token == "" {
		c.JSON(400, web.NewResponse(400, nil, "token required"))
		return false

	}
	if os.Getenv("TOKEN") != token {
		c.JSON(404, web.NewResponse(404, nil, "invalid token"))
		return false
	}
	return true
}

////////////////////////// HANDLERS //////////////////////////

func (prod *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !ValidateToken(c) {
			return
		}

		products, err := prod.service.GetAll()

		if err != nil {
			c.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v: ", err)))

		} else {
			c.JSON(200, web.NewResponse(200, products, ""))
		}

	}

}

func (prod *Product) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newProd request

		err := c.ShouldBindJSON(&newProd)

		if err != nil {
			c.String(400, "Hubo un error al cargar una persona %v: ", err)
		} else {
			response, err := prod.service.AddProduct(newProd.Name, newProd.Color, newProd.Price, newProd.Stock, newProd.Code, newProd.Published, newProd.Created)
			if err != nil {
				c.String(400, "No se pudo cargar el producto %v: ", err)
			} else {
				c.JSON(200, response)
			}
		}
	}

}

func (prod *Product) UpdateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !ValidateToken(c) {
			return
		}

		var prodToUpdate request

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(400, "El id no es válido")
		}

		err = c.ShouldBindJSON(&prodToUpdate)

		if err != nil {
			c.String(400, "Error en el body")
		} else {
			updatedProd, err := prod.service.UpdateProduct(id, prodToUpdate.Name, prodToUpdate.Color, prodToUpdate.Price, prodToUpdate.Stock, prodToUpdate.Code, prodToUpdate.Published, prodToUpdate.Created)
			if err != nil {
				c.JSON(400, err.Error())
			} else {
				c.JSON(200, updatedProd)
			}
		}

	}

}

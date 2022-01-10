package handler

import (
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/TM/proyecto/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float32 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Published    bool    `json:"published"`
	CreationDate string  `json:"creation_date"`
}

type Product struct {
	service products.Service
}

func NewProduct(ser products.Service) *Product {
	return &Product{service: ser}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.service.GetAll()

		if err != nil {
			c.String(400, "no hay nada. %v", err)
		} else {
			c.JSON(200, products)
		}
	}
}

func (p *Product) Save() gin.HandlerFunc {
	return func(c *gin.Context) {
		var prod request

		err := c.ShouldBindJSON(&prod)

		if err != nil {
			c.String(400, "error con los datos. %v", err)
		} else {
			res, err := p.service.Save(prod.Name, prod.Color, prod.Price, prod.Stock, prod.Code, prod.Published, prod.CreationDate)
			if err != nil {
				c.String(400, "no se pudo cargar el producto. %v", err)
			} else {
				c.JSON(200, res)
			}
		}
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var prod request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.String(400, "error con el id")
		}

		err = c.ShouldBindJSON(&prod)

		if err != nil {
			c.String(400, "error con el body")
		} else {
			prodUpdated, err := p.service.Update(int(id), prod.Name, prod.Color, prod.Price, prod.Stock, prod.Code, prod.Published, prod.CreationDate)
			if err != nil {
				c.JSON(400, err.Error())
			} else {
				c.JSON(200, prodUpdated)
			}
		}
	}
}

package handler

import (
	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/proyecto/internal/products"
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
			res, err := p.service.Save(prod)
			if err != nil {
				c.String(400, "no se pudo cargar el producto. %v", err)
			} else {
				c.JSON(200, res)
			}
		}
		products, err := p.service.Save()

		if err != nil {
			c.String(400, "no hay nada. %v", err)
		} else {
			c.JSON(200, products)
		}
	}
}

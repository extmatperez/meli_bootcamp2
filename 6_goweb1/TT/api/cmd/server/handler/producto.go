package handler

import (
	product "github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/internal/producto"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service product.Service
}

type Request struct {
	Name        string  `json:"name"`
	Color       string  `json:"color"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

func newProduct(ser product.Service) *Product {
	return &Product{service: ser}
}

func (pro *Product) loadFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := pro.service.loadFile()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, "El archivo fue cargado con exito")
		}
	}
}

func (pro *Product) getAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := pro.service.getAll()

		if err != nil {
			ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, products)
		}

	}
}

func (pro *Product) store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Request

		err := ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar los datos %v", err)
		} else {
			response, err := pro.service.store(product.Name, product.Color, product.Price, product.Stock, product.Code, product.IsPublished, product.CreatedAt)
			if err != nil {
				ctx.String(400, "No se pudo guardad el producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

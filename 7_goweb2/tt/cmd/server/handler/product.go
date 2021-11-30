package handler

import (
	products "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/7_goweb2/tt/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	ID          int
	Name        string
	Color       string
	Price       float64
	Stock       int
	Code        int
	IsPublished bool
	CreatedAt   string
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func (prod *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := prod.service.GetAll()
		if err != nil {
			ctx.String(400, "hubo un error: %v", err)
		} else {
			ctx.JSON(200, products)
		}
	}
}

func (prod *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product request
		err := ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.String(400, "hubo un error: %v", err)
		} else {
			response, err := prod.service.Store(product.ID, product.Name, product.Color, product.Price, product.Stock, product.Code, product.IsPublished, product.CreatedAt)
			if err != nil {
				ctx.String(400, "hubo un error al cargar el producto %v", err)
			}
			ctx.JSON(200, response)
		}
	}
}

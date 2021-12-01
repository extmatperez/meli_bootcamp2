package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lean1097/meli_bootcamp2/7_goweb2/TT/internal"
)

type request struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type Product struct {
	service internal.Service
}

func NewProduct(prod internal.Service) *Product {
	return &Product{
		service: prod,
	}
}

func (prod *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		product, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, product)
	}
}

func (prod *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := prod.service.Store(req.Name, req.Price, req.Stock)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, product)
	}
}

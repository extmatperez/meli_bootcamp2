package handler

import (
	"os"
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/8_goweb3/tm/internal/products"
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

func validateToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	envToken := os.Getenv("TOKEN")
	if token == "" {
		ctx.String(400, "falta token")
		return false
	}
	if token != envToken {
		ctx.String(400, "token invalido")
		return false
	}
	return true
}

func (prod *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
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
		if !validateToken(ctx) {
			return
		}
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

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "el id no es válido"})
			return
		}
		var prod request
		err = ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if prod.Name == "" {
			ctx.JSON(400, gin.H{"error": "el nombre no puede estar vacío"})
			return
		}
		if prod.Color == "" {
			ctx.JSON(400, gin.H{"error": "el color no puede estar vacío"})
			return
		}
		if prod.Price == 0 {
			ctx.JSON(400, gin.H{"error": "el precio no puede estar vacío"})
			return
		}
		if prod.Stock == 0 {
			ctx.JSON(400, gin.H{"error": "el stock no puede estar vacío"})
			return
		}
		if prod.Code == 0 {
			ctx.JSON(400, gin.H{"error": "el código no puede estar vacío"})
			return
		}
		if !prod.IsPublished {
			ctx.JSON(400, gin.H{"error": "el valor no puede estar vacío"})
			return
		}
		if prod.CreatedAt == "" {
			ctx.JSON(400, gin.H{"error": "la fecha no puede estar vacía"})
			return
		}
		p, err := p.service.Update(int(id), prod.Name, prod.Color, prod.Price, prod.Stock, prod.Code, prod.IsPublished, prod.CreatedAt)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

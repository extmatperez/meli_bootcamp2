package handler

import (
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/TT/internal/products"
	"github.com/gin-gonic/gin"
)

var TOKEN_PRODUCTS string = "TOKEN-PRODUCTS"

type request struct {
	Name       string `json:"name"`
	Color      string `json:"color"`
	Stock      int    `json:"stock"`
	Code       string `json:"code"`
	Published  bool   `json:"published"`
	Created_at string `json:"created_at"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{
		service: s,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidated, code, message := validateToken(ctx.GetHeader("token"))

		if !tokenValidated {
			ctx.JSON(code, gin.H{
				"message": message,
			})
			return
		}

		products, err := p.service.GetAll()

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		queryParamsAux := ctx.Request.URL.Query()
		var queryParams = map[string]string{}

		for key, val := range queryParamsAux {
			queryParams[key] = val[0]
		}

		products = p.service.FilterProducts(products, queryParams)

		ctx.JSON(200, gin.H{
			"products": products,
		})
	}
}

func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidated, code, message := validateToken(ctx.GetHeader("token"))

		if !tokenValidated {
			ctx.JSON(code, gin.H{
				"message": message,
			})
			return
		}

		var productRequest request

		err := ctx.ShouldBindJSON(&productRequest)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := p.service.Store(productRequest.Name, productRequest.Color, productRequest.Stock, productRequest.Code, productRequest.Published, productRequest.Created_at)

		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(201, gin.H{
			"product": product,
		})
	}
}

func (p *Product) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidated, code, message := validateToken(ctx.GetHeader("token"))

		if !tokenValidated {
			ctx.JSON(code, gin.H{
				"message": message,
			})
			return
		}

		productId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": "ID invalido",
			})
			return
		}

		product, err := p.service.FindById(productId)

		if err != nil {
			ctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"product": product,
		})

	}
}

func validateToken(tokenHeader string) (bool, int, string) {
	if tokenHeader == "" {
		return false, 400, "Missing token"
	}

	if tokenHeader != TOKEN_PRODUCTS {
		return false, 401, "Don´t have permission to access"
	}

	return true, 0, ""
}

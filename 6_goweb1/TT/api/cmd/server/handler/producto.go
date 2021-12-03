package handler

import (
	"fmt"
	"os"
	"strconv"

	product "github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/internal/producto"
	"github.com/extmatperez/meli_bootcamp2/tree/panceri_santiago/6_goweb1/TT/api/pkg/web"
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

func NewProduct(ser product.Service) *Product {
	return &Product{service: ser}
}

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {

		ctx.JSON(400, web.NewResponse(400, nil, "Falta token"))
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.JSON(404, web.NewResponse(404, nil, "Token Incorrecto"))

		return false
	}

	return true
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (pro *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		products, err := pro.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))

		} else {
			ctx.JSON(200, web.NewResponse(200, products, ""))
		}

	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body Request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (pro *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var product Request

		err := ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar los datos %v", err)))
		} else {
			response, err := pro.service.Store(product.Name, product.Color, product.Price, product.Stock, product.Code, product.IsPublished, product.CreatedAt)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar los datos %v", err)))
			} else {
				ctx.JSON(200, web.NewResponse(200, response, ""))
			}
		}
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body Request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [delete]
func (pro *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(200, web.NewResponse(200, nil, fmt.Sprintf("Hubo un error: %v", err)))

			ctx.String(400, "Hubo un error")
			return
		}

		mes, err := pro.service.Delete(int64(id))

		if err != nil {
			ctx.JSON(200, web.NewResponse(200, nil, fmt.Sprintf("No se pudo guardad el producto %v: ", err)))

			ctx.String(400, "No se pudo guardad el producto %v: ", err)
		} else {
			ctx.JSON(200, web.NewResponse(200, mes, ""))

		}
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body Request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [put]
func (pro *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var product Request
		idStr := ctx.Param("ID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(400, "Hubo un error")
			return
		}

		err = ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar los datos %v", err)
		} else {
			response, err := pro.service.Update(int64(id), product.Name, product.Color, product.Price, product.Stock, product.Code, product.IsPublished, product.CreatedAt)
			if err != nil {
				ctx.String(400, "No se pudo actualizar el producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body Request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [patch]
func (pro *Product) UpdateNombre() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !validarToken(ctx) {
			return
		}

		var product Request
		idStr := ctx.Param("ID")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.String(400, "Hubo un error")
			return
		}

		err = ctx.ShouldBindJSON(&product)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar los datos %v", err)
		} else {
			response, err := pro.service.UpdateNombre(int64(id), product.Name)
			if err != nil {
				ctx.String(400, "No se pudo actualizar el nombre del producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}
	}
}

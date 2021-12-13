package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/11_testing2/tm/pkg/web"
	products "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/11_testing2/tt/internal/product"

	"github.com/gin-gonic/gin"
)

type request struct {
	Color  string  `json:"color"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

type Product struct {
	service products.Service
}

func NewProduct(ser products.Service) *Product {
	return &Product{service: ser}
}

func validarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "Falta token"))
		// ctx.String(400, "Falta token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.JSON(404, web.NewResponse(404, nil, "Token incorrecto"))
		// ctx.String(404, "Token incorrecto")
		return false
	}

	return true
}

// ListProducts godoc
// @Summary List products
// @Tags Product
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/get [get]
func (prod *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// if !validarToken(ctx) {
		// 	return
		// }

		products, err := prod.service.GetAll()

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			// ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, web.NewResponse(200, products, ""))
			// ctx.JSON(200, products)
		}
	}
}

// StoreProducts godoc
// @Summary Store product
// @Tags Product
// @Description store product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products/add [post]
func (controller *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// if !validarToken(ctx) {
		// 	return
		// }

		var prod request

		err := ctx.ShouldBindJSON(&prod)

		if err != nil {
			ctx.String(400, "Hubo un error al querer cargar un producto %v", err)
		} else {
			response, err := controller.service.Store(prod.Color, prod.Price, prod.Amount)
			if err != nil {
				ctx.String(400, "No se pudo cargar la persona %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (controller *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// if !validarToken(ctx) {
		// 	return
		// }

		var prod request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&prod)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			updatedProduct, err := controller.service.Update(int(id), prod.Color, prod.Price, prod.Amount)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, updatedProduct)
			}
		}

	}
}

func (controller *Product) UpdatePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// if !validarToken(ctx) {
		// 	return
		// }

		var prod request

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = ctx.ShouldBindJSON(&prod)

		if err != nil {
			ctx.String(400, "Error en el body")
		} else {
			if prod.Price == 0 {
				ctx.String(404, "El precio no puede estar vacío")
				return
			}
			updatedProduct, err := controller.service.UpdatePrice(int(id), prod.Price)
			if err != nil {
				ctx.JSON(400, err.Error())
			} else {
				ctx.JSON(200, updatedProduct)
			}
		}

	}
}

func (controller *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// if !validarToken(ctx) {
		// 	return
		// }

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.String(400, "El id es invalido")
		}

		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, err.Error())
		} else {
			ctx.String(200, "El producto %d ha sido eliminada", id)
		}

	}
}

package handler

import (
	"os"
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/8_goweb3/tm/internal/products"
	web "github.com/extmatperez/meli_bootcamp2/tree/ottavianelli_rocio/meli_bootcamp2/8_goweb3/tm/pkg/web"
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
		ctx.JSON(400, web.NewResponse(400, nil, "token vacío"))
		return false
	}
	if token != envToken {
		ctx.JSON(400, web.NewResponse(400, nil, "token inválido"))
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
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewResponse(200, products, ""))
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
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			response, err := prod.service.Store(product.ID, product.Name, product.Color, product.Price, product.Stock, product.Code, product.IsPublished, product.CreatedAt)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			}
			ctx.JSON(200, web.NewResponse(200, response, ""))
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
			ctx.JSON(400, web.NewResponse(400, nil, "id inválido"))
			return
		}
		var prod request
		err = ctx.ShouldBindJSON(&prod)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		if prod.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el nombre no puede estar vacío"))
			return
		}
		if prod.Color == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "el color no puede estar vacío"))
			return
		}
		if prod.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "el precio no puede estar vacío"))
			return
		}
		if prod.Stock == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "el stock no puede estar vacío"))
			return
		}
		if prod.Code == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "el código no puede estar vacío"))
			return
		}
		if !prod.IsPublished {
			ctx.JSON(400, web.NewResponse(400, nil, "el valor no puede estar vacío"))
			return
		}
		if prod.CreatedAt == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "la fecha no puede estar vacía"))
			return
		}
		p, err := p.service.Update(int(id), prod.Name, prod.Color, prod.Price, prod.Stock, prod.Code, prod.IsPublished, prod.CreatedAt)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "id inválido"))
			return
		}
		err = p.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		}
		ctx.JSON(200, web.NewResponse(200, id, "producto eliminado"))
	}
}

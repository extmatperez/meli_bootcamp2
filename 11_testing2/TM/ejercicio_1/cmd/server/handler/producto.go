package handler

import (
	"fmt"
	"os"
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/tree/parra_diego/11_testing2/TM/ejercicio_1/internal/productos"
	"github.com/extmatperez/meli_bootcamp2/tree/parra_diego/8_goweb3/ejercicio_1_swagger/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	// Id      int     `json:"id"`
	Name    string  `json:"name"`
	Color   string  `json:"color"`
	Price   float64 `json:"price"`
	Stock   int     `json:"stock"`
	Code    string  `json:"code"`
	Publish bool    `json:"publish"`
	Date    string  `json:"date"`
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
		ctx.JSON(400, web.NewResponse(400, nil, "No se recibio el token"))
		//ctx.String(400, "No se recibio el token")
		return false
	}
	tokenENV := os.Getenv("TOKEN")
	if token != tokenENV {
		ctx.JSON(400, web.NewResponse(404, nil, "Token incorrecto"))
		//ctx.String(404, "Token incorrecto")
		return false
	}

	return true
}

// products godoc
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
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			// ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, products)
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
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]

func (controller *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		var produ request

		err := ctx.ShouldBindJSON(&produ)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			// ctx.String(400, "Hubo un error al querer cargar un producto %v", err)
		} else {
			response, err := controller.service.Store(produ.Name, produ.Color, produ.Price, produ.Stock, produ.Code, produ.Publish, produ.Date)
			if err != nil {
				ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
				// ctx.String(400, "No se pudo cargar un producto %v", err)
			} else {
				ctx.JSON(200, response)
			}
		}

	}
}

func (controller *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}

		var pro request
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, "Id incorrecto"))
			// ctx.String(404, "Id incorrecto")

		}
		err = ctx.ShouldBindJSON(&pro)

		if pro.Name == "" || pro.Color == "" || pro.Price == 0 || pro.Stock == 0 || pro.Code == "" || pro.Date == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "Todos los campos son requeridos"))
			// ctx.String(400, "Todos los campos son requeridos")
		} else {

			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, "Error en el body"))
				// ctx.String(400, "Error en el body")
			} else {
				produUpdate, err := controller.service.Update(int(id), pro.Name, pro.Color, pro.Price, pro.Stock, pro.Code, pro.Publish, pro.Date)
				if err != nil {
					ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
					// ctx.JSON(400, err.Error())
				} else {
					ctx.JSON(200, produUpdate)
				}
			}
		}
	}
}

func (controller *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			// ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			// ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"data": fmt.Sprintf("el producto %d fue eliminado", id)})
	}
}

func (controller *Product) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validarToken(ctx) {
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "id invalido"))
			// ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			// ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" || req.Price == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "Nombre y precio requerido"))
			// ctx.JSON(404, gin.H{"error": "Nombre y precio requerido"})
			return
		}
		p, err := controller.service.UpdateNamePrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			// ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)

	}
}

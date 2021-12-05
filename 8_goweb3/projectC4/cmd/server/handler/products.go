package handler

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	products "github.com/extmatperez/meli_bootcamp2/8_goweb3/projectC4/internal/products"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/projectC4/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

type Product struct {
	service products.Service
}

func NewProduct(prod products.Service) *Product {
	return &Product{
		service: prod,
	}
}

func validation(req request) string {
	reqValue := reflect.ValueOf(req)
	for i := 0; i < reqValue.NumField(); i++ {
		value := reqValue.Field(i).Interface()
		tipe := reflect.TypeOf(value).Kind()

		if fmt.Sprint(tipe) == "string" {
			if value == "" {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "int64" {
			if value.(int64) == 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		}

	}
	return ""
}

func validateToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")

	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "Falta Token"))
		return false
	}

	tokenENV := os.Getenv("TOKEN")

	if token != tokenENV {
		ctx.JSON(401, web.NewResponse(401, nil, "Token incorrecto"))
		return false
	}
	return true
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products/add [post]
func (controller *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		var prod request
		err := ctx.ShouldBindJSON(&prod)
		prodValidate := validation(prod)
		if prodValidate != "" {
			ctx.String(400, prodValidate)
			return
		}

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": "Error al cargar el producto",
			})
			return
		} else {
			response, err := controller.service.Store(prod.Name, prod.Color, prod.Price, prod.Stock, prod.Code,
				prod.Published, prod.CreationDate)

			if err != nil {
				ctx.JSON(400, gin.H{
					"error": "Error al crear el producto",
				})
				return
			} else {
				ctx.JSON(200, gin.H{
					"data": response,
				})
			}
		}
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products/get [get]
func (prod *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}

		product, err := prod.service.GetAll()
		if err != nil {
			ctx.JSON(401, web.NewResponse(404, nil, fmt.Sprintf("Hubo un error %v", err)))
			return
		} else {

			ctx.JSON(200, web.NewResponse(200, product, ""))
			return
		}

	}
}

func (controller *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !validateToken(ctx) {
			return
		}
		//token_header := ctx.Request.Header.Get("token")
		// token := ctx.Request.Header.Get(os.Getenv("TOKEN"))
		// if token != token_header {
		// 	ctx.JSON(401, gin.H{"error": "Token invalido"})
		// 	return
		// }
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		//id := ctx.Param("id")
		if err != nil {
			ctx.JSON(401, gin.H{"error": "ID invalido"})
			return
		}
		var prod request
		errr := ctx.ShouldBind(&prod)
		if err != nil {
			ctx.JSON(404, gin.H{"error": errr.Error()})
			return
		}

		product, errr := controller.service.Update(id, prod.Name, prod.Color, prod.Price, prod.Stock, prod.Code,
			prod.Published, prod.CreationDate)
		if errr != nil {
			ctx.JSON(404, gin.H{"error": errr.Error()})
			return
		}
		ctx.JSON(200, product)
	}
}

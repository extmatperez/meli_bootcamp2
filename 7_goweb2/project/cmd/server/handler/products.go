package handler

import (
	"fmt"
	"reflect"

	products "github.com/extmatperez/meli_bootcamp2/7_goweb2/project/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published" defautl:"false"`
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

func (controller *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

func (prod *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")

		if token != "123456" {
			ctx.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		} else {
			product, err := prod.service.GetAll()

			if err != nil {
				ctx.JSON(404, gin.H{
					"error": err,
				})
				return
			} else {
				ctx.JSON(200, gin.H{
					"data": product,
				})
				return
			}
		}

	}
}

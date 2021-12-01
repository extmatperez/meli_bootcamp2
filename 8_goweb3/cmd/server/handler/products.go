package handler

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	products "github.com/extmatperez/meli_bootcamp2/tree/ziliotto_matias/8_goweb3/internal/products"
	"github.com/gin-gonic/gin"
)

var TOKEN_PRODUCTS string = "TOKEN-PRODUCTS"

type request struct {
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      int     `json:"stock"`
	Code       string  `json:"code"`
	Published  bool    `json:"published"`
	Created_at string  `json:"created_at"`
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
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		product, err := p.service.Store(productRequest.Name, productRequest.Color, productRequest.Price, productRequest.Stock, productRequest.Code, productRequest.Published, productRequest.Created_at)

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

func (p *Product) LoadProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidated, code, message := validateToken(ctx.GetHeader("token"))

		if !tokenValidated {
			ctx.JSON(code, gin.H{
				"message": message,
			})
			return
		}

		err := p.service.LoadProducts()
		if err != nil {
			ctx.JSON(code, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(code, gin.H{
			"message": "Products loaded!",
		})
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidated, code, message := validateToken(ctx.GetHeader("token"))

		if !tokenValidated {
			ctx.JSON(code, gin.H{
				"message": message,
			})
			return
		}

		productId, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if errParse != nil {
			ctx.JSON(400, gin.H{
				"message": "ID invalido",
			})
			return
		}

		var productRequest request
		errBind := ctx.ShouldBindJSON(&productRequest)

		if errBind != nil {
			ctx.JSON(400, gin.H{
				"error": errBind.Error(),
			})
			return
		}

		var requiredFields []string
		requiredFields = append(requiredFields, "name", "color", "price", "stock", "code", "published", "created_at")

		validated, message := validateRequiredData(productRequest, requiredFields)

		if !validated {
			ctx.JSON(404, gin.H{
				"error": message,
			})
		}

		product, err := p.service.Update(productId, productRequest.Name, productRequest.Color, productRequest.Price, productRequest.Stock, productRequest.Code, productRequest.Published, productRequest.Created_at)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"product": product,
		})
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidated, code, message := validateToken(ctx.GetHeader("token"))

		if !tokenValidated {
			ctx.JSON(code, gin.H{
				"message": message,
			})
			return
		}

		productId, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if errParse != nil {
			ctx.JSON(400, gin.H{
				"message": "ID invalido",
			})
			return
		}

		err := p.service.Delete(productId)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"message": fmt.Sprintf("Product %d deleted", productId),
		})
	}
}

func (p *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenValidated, code, message := validateToken(ctx.GetHeader("token"))

		if !tokenValidated {
			ctx.JSON(code, gin.H{
				"message": message,
			})
			return
		}

		productId, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if errParse != nil {
			ctx.JSON(400, gin.H{
				"message": "ID invalido",
			})
			return
		}

		var productRequest request
		errBind := ctx.ShouldBindJSON(&productRequest)

		if errBind != nil {
			ctx.JSON(400, gin.H{
				"error": errBind.Error(),
			})
			return
		}

		var requiredFields []string
		requiredFields = append(requiredFields, "name", "price")

		validated, message := validateRequiredData(productRequest, requiredFields)

		if !validated {
			ctx.JSON(404, gin.H{
				"error": message,
			})
			return
		}

		product, err := p.service.UpdateNameAndPrice(productId, productRequest.Name, productRequest.Price)

		if err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
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

/*
	Params:
		-struct to validate
		-slice of required fields to validate
	Return:
		-bool: if everything was validated ok or not
		-string: required field error message
*/
func validateRequiredData(productRequest request, requiredFields []string) (bool, string) {
	productTypeOf := reflect.TypeOf(productRequest)

	for _, field := range requiredFields {
		fieldIndex := 0

		for fieldIndex = 0; fieldIndex < productTypeOf.NumField(); fieldIndex++ {
			if strings.ToLower(productTypeOf.Field(fieldIndex).Name) == field {
				break
			}
		}

		typeOfVField := fmt.Sprint(productTypeOf.Field(fieldIndex).Type.Kind())
		valueOfField := fmt.Sprintf("%v", reflect.ValueOf(productRequest).Field(fieldIndex).Interface())

		if !validateRequiredField(typeOfVField, valueOfField) {
			return false, "Field '" + field + "' is required"
		}
	}

	return true, ""
}

/*
	Params:
		-field type: "string", "int", etc.
		-value of field in string format: "this is a string", "2005.50", "true", etc.
	Return:
		-bool: if field was validated ok or not
*/
func validateRequiredField(fieldType, value string) bool {
	switch fieldType {
	case "string":
		if value != "" {
			return true
		}
	case "float64":
		floatVal, err := strconv.ParseFloat(value, 64)
		if err == nil && floatVal > 0 {
			return true
		}
	case "int":
		intVal, err := strconv.Atoi(value)
		if err == nil && intVal > 0 {
			return true
		}
	case "bool":
		if value != "" && (value == "true" || value == "false") {
			return true
		}
	default:
		return false
	}

	return false
}

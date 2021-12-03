package handler

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	products "github.com/extmatperez/meli_bootcamp2/9_goweb4/internal/products"
	"github.com/extmatperez/meli_bootcamp2/9_goweb4/pkg/web"
	"github.com/gin-gonic/gin"
)

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

// GetAllProducts godoc
// @Summary      Show all products
// @Description  show all products
// @Tags         Products
// @Produce      json
// @Param        token header string true "token"
// @Param        name query string false "filter products by name"
// @Param        color query string false "filter products by color"
// @Param        price query number false "filter products by price"
// @Param        stock query int false "filter products by stock"
// @Param        code query string false "filter products by code"
// @Param        published query bool false "filter products by published"
// @Param        created_at query string false "filter products by created_at"
// @Success      200  {object}  web.Response
// @Failure      400  {object}  web.Response
// @Failure      401  {object}  web.Response
// @Failure      500  {object}  web.Response
// @Router       /products/ [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.service.GetAll()

		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		queryParamsAux := ctx.Request.URL.Query()
		var queryParams = map[string]string{}

		for key, val := range queryParamsAux {
			queryParams[key] = val[0]
		}

		products = p.service.FilterProducts(products, queryParams)

		ctx.JSON(200, web.NewResponse(200, gin.H{"products": products}, ""))
	}
}

// StoreProduct godoc
// @Summary      Store Product
// @Description  store product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        token header string true "token"
// @Param        product body request true "product to store"
// @Success      200  {object}  web.Response
// @Failure      400  {object}  web.Response
// @Failure      401  {object}  web.Response
// @Failure      500  {object}  web.Response
// @Router       /products [post]
func (p *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productRequest request

		err := ctx.ShouldBindJSON(&productRequest)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		product, err := p.service.Store(productRequest.Name, productRequest.Color, productRequest.Price, productRequest.Stock, productRequest.Code, productRequest.Published, productRequest.Created_at)

		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}

		ctx.JSON(201, web.NewResponse(201, gin.H{"product": product}, ""))
	}
}

// FindProductByID godoc
// @Summary      Find Product
// @Description  find product by ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        token header string true "token"
// @Param        product_id path int true "product id"
// @Success      200  {object}  web.Response
// @Failure      400  {object}  web.Response
// @Failure      401  {object}  web.Response
// @Failure      404  {object}  web.Response
// @Router       /products/{product_id} [get]
func (p *Product) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		product, err := p.service.FindById(productId)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, gin.H{"product": product}, ""))
	}
}

// UpdateProduct godoc
// @Summary      Update Product
// @Description  update product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        token header string true "token"
// @Param        product_id path int true "product id"
// @Param        product body request true "product to update"
// @Success      200  {object}  web.Response
// @Failure      400  {object}  web.Response
// @Failure      401  {object}  web.Response
// @Failure      404  {object}  web.Response
// @Router       /products/{product_id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if errParse != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID Inválido"))
			return
		}

		var productRequest request
		errBind := ctx.ShouldBindJSON(&productRequest)

		if errBind != nil {
			ctx.JSON(400, web.NewResponse(400, nil, errBind.Error()))
			return
		}

		var requiredFields []string
		requiredFields = append(requiredFields, "Name", "Color", "Price", "Stock", "Code", "Published", "Created_at")

		validated, message := validateRequiredData(productRequest, requiredFields)

		if !validated {
			ctx.JSON(400, web.NewResponse(400, nil, message))
			return
		}

		product, err := p.service.Update(productId, productRequest.Name, productRequest.Color, productRequest.Price, productRequest.Stock, productRequest.Code, productRequest.Published, productRequest.Created_at)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, gin.H{"product": product}, ""))
	}
}

// DeleteProductByID godoc
// @Summary      Delete Product
// @Description  Delete product by ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        token header string true "token"
// @Param        product_id path int true "product id"
// @Success      200  {object}  web.Response
// @Failure      400  {object}  web.Response
// @Failure      401  {object}  web.Response
// @Failure      404  {object}  web.Response
// @Router       /products/{product_id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if errParse != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID Inválido"))
			return
		}

		err := p.service.Delete(productId)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, nil, fmt.Sprintf("Product %d deleted", productId)))
	}
}

// UpdateNameAndPriceProduct godoc
// @Summary      Update Name & Price Of Product
// @Description  update name and price of product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        token header string true "token"
// @Param        product_id path int true "product id"
// @Param        name body request true "new name of the product to update"
// @Param        price body request true "new price of the product to update"
// @Success      200  {object}  web.Response
// @Failure      400  {object}  web.Response
// @Failure      401  {object}  web.Response
// @Failure      404  {object}  web.Response
// @Router       /products/{product_id} [patch]
func (p *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productId, errParse := strconv.ParseInt(ctx.Param("id"), 10, 64)

		if errParse != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}

		var productRequest request
		errBind := ctx.ShouldBindJSON(&productRequest)

		if errBind != nil {
			ctx.JSON(400, web.NewResponse(400, nil, errBind.Error()))
			return
		}

		var requiredFields []string
		requiredFields = append(requiredFields, "Name", "Price")

		validated, message := validateRequiredData(productRequest, requiredFields)

		if !validated {
			ctx.JSON(400, web.NewResponse(400, nil, message))
			return
		}

		product, err := p.service.UpdateNameAndPrice(productId, productRequest.Name, productRequest.Price)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}

		ctx.JSON(200, web.NewResponse(200, gin.H{"product": product}, ""))
	}
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
	productValueOf := reflect.ValueOf(productRequest)

	for _, field := range requiredFields {
		fieldByName, fieldFound := productTypeOf.FieldByName(field) // Try to get the field

		if !fieldFound {
			return false, "Field '" + strings.ToLower(field) + "'not found"
		}

		// Get the type of the field
		typeOfVField := fieldByName.Type.Kind()
		// Get value of the field as string format
		valueOfField := fmt.Sprintf("%v", productValueOf.FieldByName(field).Interface())

		if !validateRequiredField(typeOfVField, valueOfField) {
			return false, "Field '" + strings.ToLower(field) + "' is required"
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
func validateRequiredField(fieldType reflect.Kind, value string) bool {
	switch fieldType {
	case reflect.String:
		if value != "" {
			return true
		}
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err == nil && floatVal > 0 {
			return true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.Atoi(value)
		if err == nil && intVal > 0 {
			return true
		}
	case reflect.Bool:
		if value != "" && (value == "true" || value == "false") {
			return true
		}
	default:
		return false
	}

	return false
}

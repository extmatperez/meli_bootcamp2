package main

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var products []Product
var token string = "TOKEN-EJ3"

type Product struct {
	Id         int64  `json:"id"`
	Name       string `json:"name" binding:"required"`
	Color      string `json:"color" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	Code       string `json:"code" binding:"required"`
	Published  bool   `json:"published" binding:"required"`
	Created_at string `json:"created_at" binding:"required"`
}

func GetAllProducts(ctx *gin.Context) {
	nameQuery := ctx.Query("name")
	colorQuery := ctx.Query("color")
	stockQuery := ctx.Query("stock")
	codeQuery := ctx.Query("code")
	publishedQuery := ctx.Query("published")
	createdAtQuery := ctx.Query("created_at")

	var productsFiltered []Product

	for i := 0; i < len(products); i++ {
		includesAllFilters := true
		if nameQuery != "" {
			if !strings.Contains(strings.ToLower(products[i].Name), strings.ToLower(nameQuery)) {
				includesAllFilters = false
			}
		}
		if colorQuery != "" {
			if !strings.Contains(strings.ToLower(products[i].Color), strings.ToLower(colorQuery)) {
				includesAllFilters = false
			}
		}
		if stockQuery != "" {
			stock, err := strconv.Atoi(stockQuery)
			if err == nil {
				if products[i].Stock != stock {
					includesAllFilters = false
				}
			}
		}
		if codeQuery != "" {
			if !strings.Contains(strings.ToLower(products[i].Code), strings.ToLower(codeQuery)) {
				includesAllFilters = false
			}
		}
		if publishedQuery != "" {
			if publishedQuery == "true" || publishedQuery == "false" {
				var published bool
				if publishedQuery == "true" {
					published = true
				} else if publishedQuery == "false" {
					published = false
				}

				if products[i].Published != published {
					includesAllFilters = false
				}
			}
		}
		if createdAtQuery != "" {
			if !strings.Contains(strings.ToLower(products[i].Created_at), strings.ToLower(createdAtQuery)) {
				includesAllFilters = false
			}
		}

		if includesAllFilters {
			productsFiltered = append(productsFiltered, products[i])
		}
	}

	ctx.JSON(200, gin.H{
		"products": productsFiltered,
	})
}

func GetProduct(ctx *gin.Context) {
	productId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "ID invalido",
		})
		return
	}

	var product Product

	for i := 0; i < len(products); i++ {
		if products[i].Id == productId {
			product = products[i]
			break
		}
	}

	if product == (Product{}) {
		ctx.JSON(404, gin.H{
			"message": "Producto no encontrado",
		})
		return
	}

	ctx.JSON(200, product)
}

func StoreProduct(ctx *gin.Context) {
	tokenHeader := ctx.GetHeader("token")
	if tokenHeader == "" {
		ctx.JSON(400, gin.H{
			"message": "Missing token",
		})
		return
	}

	if tokenHeader != token {
		ctx.JSON(401, gin.H{
			"message": "Don´t have permission to access",
		})
		return
	}

	var product Product
	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		var requiredFields []string
		requiredFields = append(requiredFields, "name", "color", "stock", "code", "published", "created_at")

		for _, fieldName := range requiredFields {
			field, validated := validateRequiredField(fieldName, strings.ToLower(err.Error()))
			if !validated {
				ctx.JSON(400, gin.H{
					"message": "Missing field '" + field + "'",
				})
				return
			}
		}

		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	product.Id = generateNewId()

	products = append(products, product)

	ctx.JSON(201, gin.H{
		"message": "Producto creado correctamente",
		"product": product,
	})
}

func validateRequiredField(fieldName, err string) (string, bool) {
	if !strings.Contains(err, "'required'") {
		return "", true
	}

	if strings.Contains(err, "'"+fieldName+"'") {
		return fieldName, false
	}

	return "", true
}

func generateNewId() int64 {
	if len(products) == 0 {
		return 1
	}

	return products[len(products)-1].Id + 1
}

func main() {
	router := gin.Default()

	products := router.Group("/products")
	{
		products.GET("/", GetAllProducts)
		products.GET("/:id", GetProduct)
		products.POST("/store", StoreProduct)
	}

	router.Run()
}

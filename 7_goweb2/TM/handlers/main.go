package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Color       string  `json:"color" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Stock       int     `json:"stock" binding:"required"`
	Code        string  `json:"code" binding:"required"`
	IsPublished bool    `json:"isPublished" binding:"required"`
	CreatedAt   string  `json:"createdAt" binding:"required"`
}

var products []Product

// var querys = []string{
// 	''
// }

func CreateNewProduct(ctx *gin.Context) {

	token := ctx.Request.Header.Get("token")

	if token != "secret_token" {
		ctx.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
		return
	}

	var lastID int = 1

	for _, prod := range products {
		if prod.ID >= lastID {
			lastID = prod.ID + 1
		}
	}

	var newProduct Product

	err := ctx.ShouldBind(&newProduct)

	if err != nil {
		newError := fmt.Sprintf("el campo %s es requerido", err)
		ctx.JSON(400, gin.H{
			"error": newError,
		})
		return
	}

	newProduct.ID = lastID

	products = append(products, newProduct)

	ctx.JSON(200, gin.H{
		"product": newProduct,
	})
}

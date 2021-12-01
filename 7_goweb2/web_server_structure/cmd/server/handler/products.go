package handler

import (
	products "github.com/extmatperez/meli_bootcamp2/pecora_estefania/7_goweb2/web_server_structure/internal/products"
	"github.com/gin-gonic/gin"
)

type request struct {
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	Code      int     `json:"code"`
	Published string  `json:"published"`
	Created   string  `json:"created"`
}

type Product struct {
	service products.Service
}

func NewProduct(serv products.Service) *Product {
	return &Product{service: serv}

}

func (controller *Product) getAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		// page, _ := strconv.ParseInt(c.Request.URL.Query().Get("page"), 10, 64)
		// limit, _ := strconv.ParseInt(c.Request.URL.Query().Get("limit"), 10, 64)

		// startIndex := (page - 1) * limit
		// endIndex := page * limit

		// var paginatedResults []Products
		// paginatedResults = prodList[startIndex:endIndex]
		token := c.GetHeader("token")
		if token != tokenPrueba {
			c.JSON(401, gin.H{
				"error": "token inválido",
			})
		} else {

			c.JSON(200, gin.H{
				"data": prodList,
			})
		}

	}

}

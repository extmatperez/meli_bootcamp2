package handler

import (
	products "github.com/extmatperez/meli_bootcamp2/pecora_estefania/8_goweb3/internal/products"
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

func (prod *Product) getAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		// token := c.GetHeader("token")
		// if token != tokenPrueba {
		// 	c.JSON(401, gin.H{
		// 		"error": "token inválido",
		// 	})
		// } else {
		products, err := prod.service.GetAll()

		if err != nil {
			c.String(400, "Hubo un error %v: ", err)
		} else {
			c.JSON(200, products)
		}

	}

	// }

}

func (prod *Product) AddProduct() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newProd request

		err := c.ShouldBindJSON(&newProd)

		if err != nil {
			c.String(400, "Hubo un error al cargar una persona %v: ", err)
		} else {
			response, err := prod.service.AddProduct(newProd.Name, newProd.Color, newProd.Price, newProd.Stock, newProd.Code, newProd.Published, newProd.Created)
			if err != nil {
				c.String("No se pudo cargar la persona %v: ", err)
			} else {
				c.JSON(200, response)
			}
		}

		// validRequest := validateKeys(req)
		// if validRequest != "" {
		// 	c.JSON(400, validRequest)
		// 	return
		// } else {
		// 	lastID++
		// 	req.ID = lastID
		// 	prodList = append(prodList, req)
		// 	c.JSON(200, req)

		// }
	}

}

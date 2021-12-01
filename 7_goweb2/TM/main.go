package main

import (
	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id             int     `json: "id"`
	Nombre         string  `json: "nombre"`
	Color          string  `json: "color"`
	Precio         float64 `json: "precio"`
	Stock          int     `json: "stock"`
	Codigo         string  `json: "codigo"`
	Publicado      bool    `json: "publicado"`
	Fecha_creacion string  `json: "fecha_creacion"`
}

var products []Producto
var lastID int

func create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productos Producto
		token := ctx.GetHeader("token")

		if token != "12345" {
			ctx.JSON(401, gin.H{
				"error": "Invalid token",
			})
			return
		}

		err := ctx.ShouldBindJSON(&productos)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			lastID++
			productos.Id = lastID

			products = append(products, productos)

			ctx.JSON(200, productos)
		}

	}
}

func main() {
	router := gin.Default()

	pr := router.Group("/productos")
	pr.POST("/", create())

	/*router.POST("/productos", func(ctx *gin.Context) {
		var productos Producto
		err := ctx.ShouldBindJSON(&productos)

		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		productos.Id = 1

		ctx.JSON(200, productos)

	})*/

	router.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
)

var productos []Producto

func main() {
	router := gin.Default()

	products := router.Group("/products")
	products.POST("/alta", alta)

	router.Run()
}

type Producto struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Apellido      string `json:"apellido" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Edad          int    `json:"edad" binding:"required"`
	Altura        int    `json:"altura" binding:"required"`
	Activo        bool   `json:"activo" binding:"required"`
	FechaCreacion string `json:"fecha_creacion" binding:"required"`
}

func alta(c *gin.Context) {
	var producto Producto
	err := c.BindJSON(&producto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		if len(productos) == 0 {
			producto.ID = 1
		} else {
			producto.ID = productos[len(productos)-1].ID + 1
		}
		productos = append(productos, producto)
		c.JSON(200, producto)
	}

}

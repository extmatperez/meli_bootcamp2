package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"productos": Productos,
	})

}
func AddProducto(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(401, gin.H{
			"error": "Not token",
		})
		return
	}
	if token != TOKEN {
		c.JSON(401, gin.H{
			"error": "Not valid token for this action",
		})
		return
	}

	lastId := Productos[len(Productos)-1].Id
	var p Producto
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p.Id = lastId + 1
	Productos = append(Productos, p)
	c.JSON(200, p)

}

// func validarCampos(productoChequear Producto) error {
// 	tipos := reflect.TypeOf(productoChequear)
// }

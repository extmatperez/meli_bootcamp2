package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
type Producto struct {
	Id             int     `json:"id"`
	Nombre         string  `json:"nombre"`
	Color          string  `json:"color"`
	Precio         float64 `json:"precio"`
	Stock          int     `json:"stock"`
	Codigo         string  `json:"codigo"`
	Publicado      bool    `json:"publicado"`
	Fecha_creacion string  `json:"fecha_creacion"`
}

func main() {

	router := gin.Default()

	router.GET("/saludo/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.JSON(http.StatusOK, gin.H{
			"msg": "hola " + nombre,
		})
	})
	router.GET("/productos", GetAll)
	router.GET("/filtro", GetFilters)

	router.Run()

}

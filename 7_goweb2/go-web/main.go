package main

import (
	"github.com/gin-gonic/gin"
)

// id, nombre, color, precio, stock, código (alfanumérico), publicado (si-no), fecha de creación.
type Producto struct {
	Id             int     `json:"id"`
	Nombre         string  `json:"nombre" binding:"required"`
	Color          string  `json:"color" binding:"required"`
	Precio         float64 `json:"precio" binding:"required"`
	Stock          int     `json:"stock" binding:"required"`
	Codigo         string  `json:"codigo" binding:"required"`
	Publicado      bool    `json:"publicado" binding:"required"`
	Fecha_creacion string  `json:"fecha_creacion" binding:"required"`
}

var Productos []Producto

const TOKEN string = "token"

func main() {

	p1 := Producto{Id: 1, Nombre: "cuchara", Color: "plata", Precio: 10.0, Stock: 40, Codigo: "eyrt47mvh4510", Publicado: true, Fecha_creacion: "29/11/2021"}
	p2 := Producto{Id: 2, Nombre: "tenedor", Color: "plata", Precio: 20.0, Stock: 5, Codigo: "765anmscid343s", Publicado: false, Fecha_creacion: "22/10/2021"}
	p3 := Producto{Id: 3, Nombre: "cuchillo", Color: "dorado", Precio: 45.0, Stock: 110, Codigo: "xmbch7cmdjgl4", Publicado: true, Fecha_creacion: "16/4/2021"}
	Productos = append(Productos, p1, p2, p3)

	router := gin.Default()
	router.GET("/productos", GetAll)
	router.POST("/productos", AddProducto)

	router.Run()

}

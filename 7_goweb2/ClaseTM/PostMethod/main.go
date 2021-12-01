package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Fecha struct {
	Dia  int
	Mes  int
	Anio int
}
type Producto struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion Fecha   `json:"fecha_creacion"`
}

func getWithId(c *gin.Context) {
	idToSearch := c.Param("id")

	data, err := os.ReadFile("/Users/nscerca/Desktop/meli_bootcamp/meli_bootcamp2/6_goweb1/Productos.json")

	if err == nil {
		var arrProductos []Producto
		var prodFiltrados []Producto
		json.Unmarshal(data, &arrProductos)

		for i, item := range arrProductos {
			if idToSearch == strconv.Itoa(item.ID) {
				prodFiltrados = append(prodFiltrados, arrProductos[i])
			}
		}

		c.JSON(200, prodFiltrados)

	} else {
		c.JSON(4040, gin.H{
			"message": "No se encontraron los datos solicitados.",
		})
	}
}

func getAll(c *gin.Context) {
	data, err := os.ReadFile("/Users/nscerca/Desktop/meli_bootcamp/meli_bootcamp2/6_goweb1/Productos.json")
	var arrProductos []Producto
	if err == nil {
		json.Unmarshal(data, &arrProductos)
		c.JSON(200, arrProductos)

	} else {
		c.JSON(4040, gin.H{
			"message": "No se encontraron los datos solicitados.",
		})
	}

}

func getAllWithFilters(c *gin.Context) {
	data, err := os.ReadFile("/Users/nscerca/Desktop/meli_bootcamp/meli_bootcamp2/6_goweb1/Productos.json")
	var arrProductos []Producto
	if err == nil {
		json.Unmarshal(data, &arrProductos)

		var prodFiltrados []Producto
		filtroNombre := c.Query("nombre")
		filtroColor := c.Query("color")
		filtroPublicado := c.Query("publicado")

		for i, item := range arrProductos {
			if item.Nombre == filtroNombre && item.Color == filtroColor && filtroPublicado == strconv.FormatBool(item.Publicado) {
				prodFiltrados = append(prodFiltrados, arrProductos[i])
			}
		}

		fmt.Println(filtroNombre)
		fmt.Println(filtroColor)
		fmt.Println(filtroPublicado)

		c.JSON(200, prodFiltrados)

	} else {
		c.JSON(4040, gin.H{
			"message": "No se encontraron los datos solicitados.",
		})
	}

}

func postProducto(c *gin.Context) {
	data, erro := os.ReadFile("/Users/nscerca/Desktop/meli_bootcamp/meli_bootcamp2/6_goweb1/Productos.json")
	var arrProductos []Producto
	if erro == nil {
		json.Unmarshal(data, &arrProductos)

		var prod Producto
		err := c.ShouldBindJSON(&prod)
		if err == nil {
			if len(arrProductos) == 0 {
				prod.ID = 1
			} else {
				fmt.Println(arrProductos)
				fmt.Println(len(arrProductos))
				// idInt := int(arrProductos[len(arrProductos)].ID)
				// prod.ID = idInt + 1
			}
			// prod.ID = arrProductos[len(arrProductos)+1]

			arrProductos = append(arrProductos, prod)
			c.JSON(201, prod)

		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	}

}
func main() {
	router := gin.Default()
	groupProducts := router.Group("api/productos")
	{
		groupProducts.POST("/add", postProducto)

		groupProducts.GET("/", getAll)
		groupProducts.GET(":id", getWithId)
		groupProducts.GET("params", getAllWithFilters)
	}
	router.Run()
}

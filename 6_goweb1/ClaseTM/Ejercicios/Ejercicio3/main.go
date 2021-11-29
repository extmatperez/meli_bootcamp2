package main

import (
	"encoding/json"
	"os"

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

func handlerSaludar(c *gin.Context) {
	saludandoA := "Hola " + c.Param("nombre")

	c.JSON(200, gin.H{
		"message": saludandoA,
	})
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
func main() {
	router := gin.Default()

	router.GET("api/Hello/:nombre", handlerSaludar)

	router.GET("api/Productos", getAll)

	router.Run()
}

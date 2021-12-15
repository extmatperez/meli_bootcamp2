package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Stock        int    `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creationDate"`
}

var products []Product

func getFiles(ctx *gin.Context) {

	data, err := os.ReadFile("../ejm/products.json") //traemos los archivos de products.json

	if err != nil {
		ctx.JSON(400, "ERROR AL IMPORTAR LOS DATOS")
		return // return para cortar la ejecucion
	}

	err = json.Unmarshal(data, &products)

	if err != nil {
		ctx.JSON(400, "ERROR AL UNMARSHALEAR") // partimos del error, si pasa es porque funciono
		return
	}

	ctx.JSON(200, products) //trajimos los products
}

func FiltrarNombre(ctx *gin.Context) {
	recibirNombre := ctx.Param("name")

	/* 	var nombresFiltrados []Product
	   	for _, product := range products {
	   		if strings.ToLower(product.Name) == strings.ToLower(recibirNombre) {
	   			nombresFiltrados = append(nombresFiltrados, product)
	   		}
	   	}
	*/

	var nombresFiltrados []Product
	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(recibirNombre)) {
			// si product.Name CONTIENE recibirNombre ... // true
			nombresFiltrados = append(nombresFiltrados, product)
		}
	}

	// filtramosopr param. creamos nuevo slice de Product donde vamos a guardar los productos filtrados
	// luego los agregamos con append al ese slice
	ctx.JSON(200, nombresFiltrados)

}

func main() {
	router := gin.Default()
	router.GET("/products", getFiles)

	filtrar := router.Group("/filtrar")

	filtrar.GET("/nombre/:name", FiltrarNombre) // al enviar un param ya la ruta es distinta
	filtrar.GET("/nombre", FiltrarNombre)       // las rutas no se pisan
	// al [pegarle al endpoint voy a reemplazar fernet y coca por las dos variables que quiera asignarle como valores

	router.Run()

}

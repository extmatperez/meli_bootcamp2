//EJERCICIO 1 AL 3 C1-GO WEB - TM

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

func BuscarPorID(ctxt *gin.Context) {
	se := false
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var productoEncontrado Product
		for _, e := range product {
			if strconv.Itoa(e.Id) == ctxt.Param("id") {
				productoEncontrado = e
				se = true
				break
			}
		}
		if se {
			ctxt.JSON(200, productoEncontrado)
		} else {
			ctxt.String(404, "No se encontro el empleado %s", ctxt.Param("id"))
		}
	}
}

func BuscarPorColor(ctxt *gin.Context) {
	parametro := ctxt.Query("Color")
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var filtrados []*Product
		for i, v := range product {
			if parametro == v.Color {
				filtrados = append(filtrados, &product[i])
			}
		}
		if len(filtrados) == 0 {
			ctxt.String(400, "No se encontró nada")
		} else {
			ctxt.JSON(200, filtrados)
		}
	}
}

func BuscarPorPrecio(ctxt *gin.Context) {
	parametro := ctxt.Query("Precio")
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var filtrados []*Product
		for i, v := range product {
			if parametro == v.Precio {
				filtrados = append(filtrados, &product[i])
			}
		}
		if len(filtrados) == 0 {
			ctxt.String(400, "No se encontró nada")
		} else {
			ctxt.JSON(200, filtrados)
		}
	}
}

func BuscarPorPublicado(ctxt *gin.Context) {
	parametro := ctxt.Param("publicado")
	se := false
	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println("Error en lectura de archivo")
	} else {
		var product []Product
		json.Unmarshal(data, &product)
		var productoEncontrado Product
		for _, e := range product {
			if strconv.FormatBool(e.Publicado) == parametro {
				productoEncontrado = e
				se = true
				break
			}
		}
		if se {
			ctxt.JSON(200, productoEncontrado)
		} else {
			ctxt.String(404, "No se encontro por publicado %s", ctxt.Param("Id"))
		}
	}
}

func main() {

	server := gin.Default()
	server.GET("/buscarId/:id", BuscarPorID)
	server.GET("/buscarPorColor", BuscarPorColor)
	server.GET("/buscarPorPrecio", BuscarPorPrecio)
	server.GET("/buscarPorPublicado/:publicado", BuscarPorPublicado)
	server.Run()

	/*
		data, err := os.ReadFile("./products.json")
		if err != nil {
			fmt.Println("Error en lectura de archivo")
		} else {
			var product []Product
			json.Unmarshal(data, &product)
			router := gin.Default()
			router.GET("/product", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": product,
				})
			})
			router.Run()
		}

	*/
	// ACEPTAR UN PARAMETRO POR URL
	/* router.GET("/productname/:nombre", func(c *gin.Context) {
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hello %s", nombre)
	}) */
}

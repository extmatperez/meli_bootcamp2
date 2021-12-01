package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {

	productos, err := leer_productos("products.json")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"productos": productos,
	})

}
func GetFilters(c *gin.Context) {
	var etiquetas []string
	etiquetas = append(etiquetas, "nombre", "color")
	var productosFiltrados []Producto
	productosFiltrados = leer_productos("products.json")
}
func leer_productos(ruta string) ([]Producto, error) {
	var productos = []Producto{}
	file, err := os.ReadFile(ruta)

	if err != nil {
		return productos, errors.New("error al leer el archivo ")

	}

	errorr := json.Unmarshal(file, &productos)
	if errorr != nil {
		return productos, errors.New("no se pudo converitr a json")
	}
	return productos, nil

}


func filtrar(sliceProductos []Producto, campo, valor string) []Persona{
	var filtrado []Producto
	var p Persona
	tipos := reflect.TypeOf(p)
	i:= 0 
	for i= 0; i<tipos.NumField(); i++ {
		if(strings.ToLower(tipos.Field(i).Name) == campo){
			break
		}
	}
	for k,v:= range sliceProductos{
		var cadena string
		cadena = fmt.Sprintf("%v",reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor)	
			filtrado = append(filtrado, v)
		}
		
	}

	return filtrado
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	transacciones := router.Group("/transacciones")
	{
		transacciones.GET("/filtros", filterTrans)
		transacciones.GET("/find/:id", getById)
		transacciones.GET("/", getAll)
	}
	saludar := router.Group("/saludar")
	{
		saludar.GET("/:name", saludo)
		saludar.GET("/", saludo)
	}

	router.Run()

}

func String(file string) {
	panic("unimplemented")
}

// estructura

type transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaCreacion     string  `json:"fecha_creacion"`
}

var transacciones []transaccion

//

func saludo(c *gin.Context) {
	queryName := c.Query("name")
	paramName := c.Param("name")

	if queryName != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola " + queryName,
		})
	} else if paramName != "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola " + paramName,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola, ingresa tu nombre como query o param",
		})
	}

}

func getById(c *gin.Context) {
	filterId := c.Param("id")
	var filtId transaccion

	if filterId != "" {
		for _, v := range transacciones {
			if filterId == strconv.Itoa(v.ID) {
				filtId = v
				break
			}
		}

	}

	if filtId.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"transacciones": filtId,
		})

	} else {
		c.JSON(http.StatusBadRequest, "el id que buscas no se encuentra")
	}

}

func getAll(c *gin.Context) {
	datos, err := os.ReadFile("./transaccion.json")
	if err != nil {
		c.String(400, "no se puede abrir el archivo")
	} else {
		json.Unmarshal(datos, &transacciones)
		c.JSON(http.StatusOK, gin.H{
			"transacciones": transacciones,
		})

	}
}

func filtrar(sliceTransacciones []transaccion, campo string, valor string) []transaccion {
	var filtrado []transaccion
	var trans transaccion

	tipos := reflect.TypeOf(trans)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		if strings.ToLower(tipos.Field(i).Name) == campo {
			break
		}
	}
	for _, v := range sliceTransacciones {
		var cadena string
		cadena = fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor) {
			filtrado = append(filtrado, v)
		}

	}

	return filtrado

}

// los filtros andan todos menos con fecha_creacion y codigo_transaccion . y tira un error de que no puede importar C

func filterTrans(c *gin.Context) {
	var etiquetas []string
	etiquetas = append(etiquetas, "moneda", "emisor", "receptor", "fecha_creacion", "codigo_transaccion")
	var transaccionesFiltradas []transaccion

	transaccionesFiltradas = transacciones

	for _, v := range etiquetas {
		if len(c.Query(v)) != 0 && len(transaccionesFiltradas) != 0 {
			transaccionesFiltradas = filtrar(transacciones, v, c.Query(v))

		}
	}

	if len(transaccionesFiltradas) == 0 {
		c.String(200, "no hay coincidencias")
	} else {
		c.JSON(200, transaccionesFiltradas)
	}

}

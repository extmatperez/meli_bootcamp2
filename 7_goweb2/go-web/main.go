package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func main() {

	router := gin.Default()
	transacciones := router.Group("/transacciones")

	transacciones.GET("/filtros", filterTrans)
	transacciones.GET("/find/:id", getById)
	transacciones.GET("/", load)
	transacciones.GET("/get", getTrans)
	transacciones.POST("/add", addTransaccion)

	saludar := router.Group("/saludar")

	saludar.GET("/:name", saludo)
	saludar.GET("/", saludo)

	router.Run()

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

// ==========================================================GET======================================================//
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

func load(c *gin.Context) {
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
		c.String(http.StatusOK, "no hay coincidencias")
	} else {
		c.JSON(http.StatusOK, transaccionesFiltradas)
	}

}

func getTrans(c *gin.Context) {

	token := c.GetHeader("token")
	if token != "" {
		if token == "39470939" {
			if len(transacciones) > 0 {
				c.JSON(http.StatusOK, transacciones)
			} else {
				c.String(http.StatusOK, "no hay transacciones cargadas")
			}

		} else {
			c.String(401, "token incorrecto")
		}

	} else {
		c.String(http.StatusBadGateway, " no ingreso token ")
	}
}

// ==========================================================POST======================================================//

func addTransaccion(c *gin.Context) {
	var trans transaccion
	token := c.GetHeader("token")

	if token != "" {
		if token == "39470939" {
			err := c.ShouldBindJSON(&trans)
			if err != nil {
				c.String(400, "se produjo un error: %v", err.Error())
				return
			} else {
				switch {
				case trans.Monto == 0.0:
					c.String(http.StatusBadGateway, " no puede el monto vacio")
				case trans.Emisor == "":
					c.String(http.StatusBadGateway, " no puede estar el emisor vacio")
				case trans.Moneda == "":
					c.String(http.StatusBadGateway, " no puede dejar de especificar la moneda")
				case trans.Receptor == "":
					c.String(http.StatusBadGateway, " no puede estar el receptor vacio")
				default:
					if len(transacciones) == 0 {
						trans.ID = 1
					} else {
						trans.ID = transacciones[len(transacciones)-1].ID + 1
					}
					currentTime := time.Now().Format("02-01-2006")
					trans.FechaCreacion = fmt.Sprintf(currentTime)
					trans.CodigoTransaccion = uuid.NewV4().String()
					transacciones = append(transacciones, trans)
					c.JSON(http.StatusOK, trans)
				}
			}

		} else {
			c.String(401, "token incorrecto")
		}

	} else {
		c.String(http.StatusBadGateway, " no ingreso token ")
	}

}

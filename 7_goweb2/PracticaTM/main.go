package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Fecha struct {
	Dia, Mes, Anio int
}
type Transaccion struct {
	Id             int     `json:"id"`
	CodTransaccion string  `json:"cod_transaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         string  `json:"emisor"`
	Receptor       string  `json:"receptor"`
	FechaTrans     string  `json:"fecha_trans"`
}

func buscarTransaccion(ctx *gin.Context) {
	var transac Transaccion

	parametro := ctx.Param("id")
	se := false
	for _, valor := range transacciones {
		if strconv.Itoa(valor.Id) == parametro {
			transac = valor
			se = true
			break
		}
	}

	if se {
		ctx.JSON(200, transac)
	} else {
		ctx.String(404, "No se encontro la transacción %s", parametro)
	}

	// if c.BindJSON(&transac) == nil {
	// 	log.Println("Bind por JSON")
	// 	log.Println("ID de transaccion: ", transac.Id)
	// 	log.Println("Codigo de transaccion: ", transac.CodTransaccion)
	// 	c.String(http.StatusOK, "(Query JSON) - Transaccion: %s, ID: %s\n", transac.CodTransaccion, transac.Id)
	// } else {
	// 	c.String(404, "La transaccion no existe")
	// }
}

func filtrarTransacciones(ctx *gin.Context) {
	//Almaceno todas las etiquetas de mi struct
	var etiquetas []string
	etiquetas = append(etiquetas, "id", "cod_transaccion", "moneda", "monto", "emisor", "receptor", "fecha_trans")

	var transacFiltradas []Transaccion

	for i, etiqueta := range etiquetas {
		if ctx.Query(etiqueta) != "" {
			//Si el valor de esa etiqueta en el GET no es vacío, lo busco
			for _, transaccion := range transacciones {
				valor := fmt.Sprintf("%v", reflect.ValueOf(transaccion).Field(i).Interface())
				if valor == ctx.Query(etiqueta) {
					transacFiltradas = append(transacFiltradas, transaccion)
				}
			}
		}
	}

	if len(transacFiltradas) == 0 {
		ctx.String(200, "No se encontró ninguna transaccion")
	} else {
		ctx.JSON(200, transacFiltradas)
	}
}

func addTransaccion(ctx *gin.Context) {
	var transac Transaccion

	err := ctx.ShouldBindJSON(&transac)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		generarIDTransaccion(&transac)
		transacciones = append(transacciones, transac)
		ctx.JSON(200, transac)
	}
}

func generarIDTransaccion(transac *Transaccion) {
	id := 0
	for _, valor := range transacciones {
		id = valor.Id
	}
	transac.Id = id + 1
}

var transacciones []Transaccion

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola Facundo",
		})
	})

	groupTransac := router.Group("/transacciones")

	//Endpoints GET
	groupTransac.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"GetAll": transacciones,
		})
	})

	groupTransac.GET("/:id", buscarTransaccion)
	groupTransac.GET("/filtros", filtrarTransacciones)

	//Endpoints POST
	groupTransac.POST("/cargar", addTransaccion)

	router.Run()
}

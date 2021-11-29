package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

///////////////////////////////////////////////////////////////////////////
type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion int     `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}
///////////////////////////////////////////////////////////////////////////
func GetAll(c *gin.Context){
	var lista []Transaccion
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &lista)
	c.JSON(200, lista)
}

///////////////////////////////////////////////////////////////////////////
func filtrarTransacciones(ctx *gin.Context){
	var lista []Transaccion
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &lista)
	var filtrados []*Transaccion

	for i, e := range lista{
		if ctx.Query("codigo") ==  e.Emisor{
			filtrados = append(filtrados, &lista[i])
		}
	}
	if len(filtrados) == 0{
		ctx.String(404, "No se encontraron coincidencias")
	}else{
		ctx.JSON(200, &filtrados)
	}
}
///////////////////////////////////////////////////////////////////////////
func buscarTransaccion(ctx *gin.Context){
	parametro := ctx.Param("codigo_transaccion")

	var lista []Transaccion
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &lista)

	var tran Transaccion
	se := false

	for _, v := range lista{
		str := fmt.Sprint(v.CodigoTransaccion)
		if str == parametro{
			tran = v
			se = true
			break
		}
	}
	if se{
		ctx.JSON(200, tran)
	}else{
		ctx.String(404, "Registro no encontrado")
	}
}





func main() {

	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, "¡Hola Juampi!")
	})

	router.GET("/transacciones", GetAll)
	router.GET("/filtrar",filtrarTransacciones)
	router.GET("/buscar/:codigo_transaccion",buscarTransaccion)

	router.Run()
	
}
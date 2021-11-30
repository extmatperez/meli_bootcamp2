package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/gin-gonic/gin"
)

//////////////////////////////Structs/////////////////////////////////////////////
type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion int     `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

////////////////////////////Crear Handler///////////////////////////////////////////////
func GetAll(c *gin.Context) {
	var lista []Transaccion
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &lista)
	c.JSON(200, lista)
}

//////////////////////////////Filtrar/////////////////////////////////////////////
func filtrarTransacciones(ctx *gin.Context) {
	var lista []Transaccion
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &lista)
	var filtrados []*Transaccion

	for i, e := range lista {
		if ctx.Query("codigo") == e.Emisor {
			filtrados = append(filtrados, &lista[i])
		}
	}
	if len(filtrados) == 0 {
		ctx.String(404, "No se encontraron coincidencias")
	} else {
		ctx.JSON(200, &filtrados)
	}
}

/////////////////////////////Buscar//////////////////////////////////////////////
func buscarTransaccion(ctx *gin.Context) {
	parametro := ctx.Param("codigo_transaccion")

	var lista []Transaccion
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &lista)

	var tran Transaccion
	se := false

	for _, v := range lista {
		str := fmt.Sprint(v.CodigoTransaccion)
		if str == parametro {
			tran = v
			se = true
			break
		}
	}
	if se {
		ctx.JSON(200, tran)
	} else {
		ctx.String(404, "Registro no encontrado")
	}
}

/////////////////////////////Validar con REFLECT//////////////////////////////////////////////

func validar(req Transaccion) string {

	r := reflect.ValueOf(req)

	for i := 0; i < r.NumField(); i++ {

		varValor := r.Field(i).Interface()
		s := reflect.TypeOf(varValor).Kind()

		if fmt.Sprint(s) == "string" {
			if varValor == "" {
				return fmt.Sprintf("El campo %v no puede estar vacio", r.Type().Field(i).Name)
			}
		} else {
			if varValor == 0 {
				return fmt.Sprintf("El campo %v no puede ser cero", r.Type().Field(i).Name)
			}
		}
	}

	return ""
}

/////////////////////////////Headers y POST//////////////////////////////////////////////

var nlista []Transaccion

func agregarTransaccion(ctx *gin.Context) {
	var req Transaccion
	err := ctx.ShouldBind(&req)
	token := ctx.GetHeader("token")

	req.ID = 1

	val := validar(req)

	if val != "" {
		ctx.String(400, val)
		return
	}

	if token == "secure" {
		if err != nil {
			ctx.String(400, "Ha ocurrido un error")
		} else {
			req.ID = len(nlista) + 1
			nlista = append(nlista, req)
			ctx.JSON(200, req)
		}
	} else {
		ctx.String(401, "No tiene permisos para realizar la petición realizada")
	}
}


//////////////////////////////////////////////////////////////////////////////////////////

func main() {

	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, "¡Hola Juampi!")
	})

	router.GET("/transacciones", GetAll)
	router.GET("/filtrar", filtrarTransacciones)
	router.GET("/buscar/:codigo_transaccion", buscarTransaccion)

	router.POST("/buscar/add", agregarTransaccion)

	router.Run()

}

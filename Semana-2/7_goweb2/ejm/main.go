package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	ID       int    `json:"id"`
	Nombre   string `json:"first_name"`
	Apellido string `json:"last_name"`
	Edad     string `json:"email"`
}

var personas []Persona

func GetPersonas(ctx *gin.Context) {

	token := ctx.GetHeader("token")
	// al header lo manda el front. con GetHeader yo me lo traigo y lo comparo
	//GetHeader devuelve un string
	if token != "" {
		if token == "123456" {
			if len(personas) > 0 {
				ctx.JSON(200, personas)
			} else {
				ctx.String(200, "No hay personas para mostrar")
			}

		} else {
			ctx.String(http.StatusUnauthorized, "No hay personas para mostrar")
		}
	} else {
		ctx.String(500, "No se ingreso ningun token")
	}

	/* 	if len(personas) > 0 {
	   		ctx.JSON(200, personas)
	   	} else {
	   		ctx.String(200, "No hay personas para mostrar")
	   	} */

}

func AddPersona(ctx *gin.Context) {

	var per Persona
	body := fmt.Sprintf("%v", ctx.Request.Body)
	if body == "" {
		ctx.JSON(400, gin.H{
			"Error": "El body esta vacio",
		})

		return

	}

	err := ctx.ShouldBindJSON(&per)

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})

	} else {
		if len(personas) == 0 {
			per.ID = 1
		} else {
			per.ID = personas[len(personas)-1].ID + 1 //importante loque hago aca
			// de personas me traigo LA ULTIMA con len(personas)-1 .   y de esa ultima persona me da el ID.
			// y a ese ID lo incremento en 1
		}
		personas = append(personas, per)
		ctx.JSON(200, per)
	}

}

func LoadData(ctx *gin.Context) {

	data, err := os.ReadFile("./personas.json")
	// para importar el archivo utilizamos el packete os, que nos va a traer el conjunto de jsons en bytes,
	// por lo que vamos a necesitar usar unmarshal
	// parto de data y err para luego manipular el error y unmarshalear
	if err != nil {
		ctx.String(400, "error al cargar archivo")
		return

	} else {

		json.Unmarshal(data, &personas)
		// importante inicar la direccion de memmoria, en este caso es el slice personas que creamos como variable global
		ctx.String(200, "personas cargadas")

	}

	// luego necesitamos validar si el unmarshall tuvo exito validandolo nuevamente
	if err != nil {
		ctx.String(400, "error al unmarshalear")
		return
	}

	//si tuvo exito debera cargar perfectamente
	ctx.JSON(200, personas)

}

// creo la funcion filtrar la cual va  recibir un slice de personas, mas el campo que avn a ser las etiquetas
// y va a devolver otro slice de personas que va a ser "filtradas"
// func Filtrar(slicePersonas []Persona campo string, valor string) []Persona {
// 	var filtrado []Persona

// 	var per Persona
// 	tipos := reflect.TypeOf(per)
// 	for i := 0; i < tipos.NumField(); i++ {

// 	if	strings.ToLower(tipos.NumField(i).Nombre) == campo {
// 		  break
// 	}
// 	}

// 	for k,v := range slicePersonas{
// 		reflect.ValueOf(v).Field(i).Interface() == valor
// 	}

// return filtrado
// }

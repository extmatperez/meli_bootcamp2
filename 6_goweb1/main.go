package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_creacion"`
}

var usuariosJSON []Usuario

func AddUsuario(ctx *gin.Context) {
	var usr Usuario
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		if len(usuariosJSON) == 0 {
			usr.ID = 1
		} else {
			usr.ID = usuariosJSON[len(usuariosJSON)-1].ID + 1
		}
		usuariosJSON = append(usuariosJSON, usr)
		ctx.JSON(200, usr)
	}
}

func GetUsuarios(ctx *gin.Context) {
	token := ctx.GetHeader("token")

	if token != "" {
		if token == "123456" {
			if len(usuariosJSON) > 0 {
				ctx.JSON(200, usuariosJSON)
			} else {
				ctx.String(200, "No hay personas cargadas")
			}
		} else {
			ctx.JSON(401, "Token incorrecto")
		}
	} else {
		ctx.String(400, "No se ingreso token")
	}
}

func filtrar(sliceUsuarios []Usuario, campo string, valor string) []Usuario {
	var filtrado []Usuario

	var usr Usuario
	tipos := reflect.TypeOf(usr)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		if strings.ToLower(tipos.Field(i).Name) == campo {
			break
		}
	}

	for _, v := range sliceUsuarios {
		var cadena string
		cadena = fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor) {
			filtrado = append(filtrado, v)
		}
	}
	return filtrado
}

func FiltrarUsuarios(ctx *gin.Context) {
	var etiquetas []string
	etiquetas = append(etiquetas, "nombre", "apellido")

	var usuariosFiltrados []Usuario

	usuariosFiltrados = usuariosJSON

	for _, v := range etiquetas {
		if len(ctx.Query(v)) != 0 && len(usuariosFiltrados) != 0 {
			usuariosFiltrados = filtrar(usuariosFiltrados, v, ctx.Query(v))
		}
	}

	if len(usuariosFiltrados) == 0 {
		ctx.String(200, "No hay coincidencias")
	} else {
		ctx.JSON(200, usuariosFiltrados)
	}
}

func LoadData(ctx *gin.Context) {
	data, err := os.ReadFile("./Usuarios.json")
	if err != nil {
		ctx.String(400, "No se pudo leer el archivo")
	} else {
		json.Unmarshal(data, &usuariosJSON)
		ctx.String(200, "Personas Cargadas")
	}
}

/*
func HandlerUsers(GetAll *gin.Context) {

	data, _ := os.ReadFile("./Usuarios.json")
	json.Unmarshal(data, &usuariosJSON)

	GetAll.JSON(200, gin.H{
		"message": usuariosJSON})
}

/*
func BuscarUsuario(ctx *gin.Context) {
	data, _ := os.ReadFile("./Usuarios.json")
	json.Unmarshal(data, &usuariosJSON)

	parametro := ctx.Param("id")
	var usr Usuario
	se := false
	for _, v := range usuariosJSON {
		if v.ID == parametro {
			usr = v
			se = true
			break
		}
	}

	if se {
		ctx.JSON(200, usr)
	} else {
		ctx.String(404, "No se encontro el empleado %s", parametro)
	}

}*/

func main() {
	router := gin.Default()
	usuarios := router.Group("/usuarios")

	usuarios.POST("/add", AddUsuario)
	usuarios.GET("/", GetUsuarios)
	usuarios.GET("/loadData", LoadData)
	usuarios.GET("/filtros", FiltrarUsuarios)
	//usuarios.GET("/:id", BuscarUsuario)

	router.Run(":8000")
}

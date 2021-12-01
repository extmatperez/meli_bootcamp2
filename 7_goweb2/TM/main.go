package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        bool   `json:"activo"`
	FechaCreacion string `json:"fecha_creacion"`
}

var usuarios_global []Usuario

func GetAll(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"data": usuarios_global,
	})
}

func filtrarUsuarios(ctx *gin.Context) {
	nombre := ctx.Query("nombre")
	apellido := ctx.Query("apellido")
	email := ctx.Query("email")
	altura, _ := strconv.ParseInt(ctx.Query("altura"), 0, 64)
	edad, _ := strconv.ParseInt(ctx.Query("edad"), 0, 64)
	activo, _ := strconv.ParseBool(ctx.Query("activo"))
	fecha := ctx.Query("fecha")

	var usuarios []Usuario
	var filtrados []Usuario
	data, err := os.ReadFile("./users.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	//data1 := Usuario{1, "Ida", "Tieman", "itieman0@npr.org", 82, 187, true, "06/15/2021"}
	json.Unmarshal(data, &usuarios)

	for _, user := range usuarios {
		if user.Nombre == nombre || user.Apellido == apellido || user.Email == email || user.Altura == int(altura) || user.Edad == int(edad) || user.Activo == bool(activo) || user.FechaCreacion == fecha {
			filtrados = append(filtrados, user)
		}
	}

	ctx.JSON(http.StatusOK, filtrados)

}

func consultarUsuario(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 64)

	var usuarios []Usuario
	var filtrado Usuario
	data, err := os.ReadFile("./users.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	json.Unmarshal(data, &usuarios)
	for _, user := range usuarios {
		if user.ID == int(id) {
			filtrado = user
			break
		}
	}

	ctx.JSON(http.StatusOK, filtrado)

}

func loadData(ctx *gin.Context) {
	data, err := os.ReadFile("./users.json")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

	json.Unmarshal(data, &usuarios_global)

	ctx.String(http.StatusOK, "Datos cargados!!")
}

func validarUsuario(usuario Usuario) error {
	message := ""
	var fields []string
	fields = append(fields, "ID", "Nombre", "Apellido", "Email", "Edad", "Altura", "Activo", "FechaCreacion")
	user := reflect.TypeOf(usuario)
	for i := 1; i < user.NumField(); i++ {
		nombre := user.Field(i).Name
		valor := reflect.ValueOf(usuario).FieldByName(fields[i]).Interface()
		if valor == "" || valor == 0 {
			message += "El campo " + nombre + " es requerido\n"
		}
	}
	if message == "" {
		return nil
	} else {
		return fmt.Errorf(message)
	}
}

func registrarUsuario(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	fmt.Println(token)
	if token != "12345" {
		ctx.String(401, "No tiene permisos para realizar la petición solicitada.")
	} else {
		var user Usuario
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println("----", err)
		} else {
			fmt.Println(len(usuarios_global))
			if len(usuarios_global) == 0 {
				user.ID = 1
			} else {
				user.ID = usuarios_global[len(usuarios_global)-1].ID + 1
			}
			err := validarUsuario(user)
			if err != nil {
				ctx.String(200, err.Error())
			} else {
				usuarios_global = append(usuarios_global, user)
			}

			ctx.JSON(http.StatusOK, user)
		}
	}

}

func main() {
	router := gin.Default()

	router.GET("/hello-world/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name)
		c.JSON(200, gin.H{
			"message": "Hola " + name + "!!",
		})
	})
	usuarios := router.Group("/usuarios")
	{
		usuarios.GET("/", GetAll)
		usuarios.GET("/:id", consultarUsuario)
		usuarios.POST("/registrar", registrarUsuario)
	}

	router.GET("/filtrarUsuarios", filtrarUsuarios)
	router.GET("loadData/", loadData)

	router.Run()

}

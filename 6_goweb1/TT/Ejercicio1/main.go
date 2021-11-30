package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func GetAll(c *gin.Context) {
	var users []Usuario
	data, err := os.ReadFile("./users.json")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	//data1 := Usuario{1, "Ida", "Tieman", "itieman0@npr.org", 82, 187, true, "06/15/2021"}

	json.Unmarshal(data, &users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
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

func main() {
	router := gin.Default()

	router.GET("/hello-world/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name)
		c.JSON(200, gin.H{
			"message": "Hola " + name + "!!",
		})
	})

	router.GET("/filtrarUsuarios", filtrarUsuarios)
	router.Run()

}

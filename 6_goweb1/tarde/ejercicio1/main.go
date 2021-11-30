package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Edad          int    `json:"edad"`
	Altura        int    `json:"altura"`
	Activo        string `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

func main() {

	router := gin.Default()
	router.GET("/users/:val", filtrado)

	router.Run()

}

func filtrado(c *gin.Context) {
	val := c.Param("val")
	var usersList []User
	var filtrados []User

	dbusers, err := ioutil.ReadFile("users.json")
	json.Unmarshal(dbusers, &usersList)

	if err != nil {
		c.String(http.StatusBadRequest, "Algo salió mal")
	} else {
		for _, u := range usersList {
			switch val {
			case strconv.Itoa(u.ID):
				filtrados = append(filtrados, u)
			case u.Nombre:
				filtrados = append(filtrados, u)
			case u.Apellido:
				filtrados = append(filtrados, u)
			case u.Email:
				filtrados = append(filtrados, u)
			case strconv.Itoa(u.Edad):
				filtrados = append(filtrados, u)
			case strconv.Itoa(u.Altura):
				filtrados = append(filtrados, u)
			case u.Activo:
				filtrados = append(filtrados, u)
			case u.FechaCreacion:
				filtrados = append(filtrados, u)
			}
		}
		if len(filtrados) > 0 {
			c.JSON(http.StatusOK, filtrados)
		} else {
			c.String(http.StatusNotFound, "No se encontro el user")
		}
	}

}

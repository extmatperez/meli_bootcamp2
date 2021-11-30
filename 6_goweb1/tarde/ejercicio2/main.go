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
	router.GET("/users/:id", filtraId)

	router.Run()

}

func filtraId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var usersList []User
	var filtrados []User

	dbusers, err := ioutil.ReadFile("users.json")
	json.Unmarshal(dbusers, &usersList)

	if err != nil {
		c.String(http.StatusBadRequest, "Algo salió mal")
	} else {
		for _, u := range usersList {
			if id == u.ID {
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

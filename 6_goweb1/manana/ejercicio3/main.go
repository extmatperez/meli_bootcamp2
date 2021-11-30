package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Users struct {
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
	router.GET("/users", GetAll)

	router.Run()

}

func GetAll(c *gin.Context) {

	var userArr []Users
	readUsers, err := os.ReadFile("./users.json")

	if err != nil {
		fmt.Println(err)
	} else {
		json.Unmarshal(readUsers, &userArr)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": userArr,
	})

}

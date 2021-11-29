package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Personas struct {
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

	var prodSalida []Personas
	prodSalida = read()
	var salida = ""
	fmt.Println("-------", prodSalida)
	for i := 0; i < len(prodSalida); i++ {
		salida += string(prodSalida[i].ID)
	}
	s := gin.New()
	s.GET("/personas", func(c *gin.Context) {
		c.String(http.StatusOK, salida)
	})

	s.Run()

}

func read() []Personas {
	data, _ := os.ReadFile("./personas.json")
	var prodSalida []Personas

	json.Unmarshal([]byte(string(data)), &prodSalida)

	fmt.Println(prodSalida)
	return prodSalida
}

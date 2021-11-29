package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

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

	fmt.Println(users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
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

	router.GET("/usuarios", GetAll)

	router.Run()

}

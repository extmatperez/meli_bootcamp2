package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shyandsy/ShyGinErrors"
)

var requestErrorMessage = map[string]string {
    "error_invalid_nombre" : "el nombre es requerido",
    "error_invalid_apellido" : "el apellido es requerido",
    "error_invalid_password" : "la edad es requerida",
}

type Persona struct {
	ID 	     int `json:"id"`
	Nombre   string `json:"nombre" binding:"required" msg:"error_invalid_nombre"`
	Apellido string `json:"apellido" binding:"required" msg:"error_invalid_apellido"`
	Edad     int `json:"edad" binding:"required" msg:"error_invalid_password"`
}


var personas []Persona

func addPersona(c *gin.Context) {
	var per Persona
	err := c.ShouldBindJSON(&per)

	if err != nil {
		ge := ShyGinErrors.NewShyGinErrors(requestErrorMessage)
		req := Persona{}
		errors := ge.ListAllErrors(req, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}
	
	if len(personas) > 0 {
		per.ID = personas[len(personas)-1].ID + 1
	} else {
		per.ID = 1
	}
	personas = append(personas, per)
	c.JSON(http.StatusCreated, per)
}

func tokenValid(tk string) bool {
    switch tk {
    case
        "123456",
        "444444",
        "555555":
        return true
    }
    return false
}

func getPersonas(c *gin.Context){
	token := c.GetHeader("token")

	if tokenValid(token) {
		if len(personas) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message":"no people found",
			})
			return
		}
		
		c.JSON(http.StatusOK, personas)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"message":"no tiene permisos para realizar la petición solicitada, por favor envíe un token correcto",
	})
}

func main() {
	r := gin.Default()

	personasEP := r.Group("/personas")
	{	
		personasEP.POST("/add", addPersona)
		personasEP.GET("/", getPersonas)
	}

	r.Run()
}
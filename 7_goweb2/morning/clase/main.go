package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Persona struct {
	ID 	     int `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int `json:"edad"`
}

var personas []Persona

func addPersona(c *gin.Context) {
	var per Persona

	err := c.ShouldBindJSON(&per)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	
	if len(personas) == 0 {
		per.ID = 1
	} else {
		per.ID = personas[len(personas)-1].ID + 1
	}
	//per.ID = len(personas) + 1
	personas = append(personas, per)
	c.JSON(http.StatusCreated, per)
}

func getPersonas(c *gin.Context){

	token := c.GetHeader("token")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":"token not found",
		})
		return
	}

	if token == "123456" {
		if len(personas) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message":"no people found",
			})
			return
		}
		
		c.JSON(http.StatusOK, personas)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"message":"incorrect token",
	})

}


func main() {
	r := gin.Default()

	personasEP := r.Group("/personas")
	{
		personasEP.GET("/", getPersonas)
		personasEP.POST("/add", addPersona)
	}

	r.Run()
}
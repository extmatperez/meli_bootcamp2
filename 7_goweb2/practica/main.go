package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
var us []Users
type Users struct {
	ID           int     `json:"id"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Active bool    `json:"active"`
	CreationDate time.Time  `json:"creation_date"`
}
func main() {
	// Routers
	router := gin.New()
	router.GET("/users", GetALL)
	router.POST("/users/add", addUser)
	router.Run()
}
func addUser(c *gin.Context) {
	// declare vars
	var (
		usRequest []Users
	)
	err := c.BindJSON(&usRequest)
	token := c.GetHeader("token")
	//authorization exception
	if token != "123456789" {
		c.String(http.StatusUnauthorized, "No tiene acceso a esta funcion")
		return
	}
	// filter exception
	field, errField := checkerData(usRequest)

	if errField {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "El campo "+field+" Es requerido",
		})
		return
	}
	//error exception and creation
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		for _, v := range usRequest {
			v.ID = len(us)+1
			v.Active = true
			v.CreationDate = time.Now()
			us = append(us, v)
		}
		c.String(http.StatusCreated, "Creado ok")
	}
}
func GetALL(c *gin.Context) {
	c.JSON(200, us)
}

func checkerData (u []Users) (string, bool) {
	for _, v := range u {
		if v.FirstName == "" {
			return "nombre", true
		}
		if v.LastName == "" {
			return "apellido", true
		}
		if v.Age < 0 {
			return "Edad", true
		}
		if v.Email == ""  {
			return "Email", true
		}
		if v.Height == 0  {
			return "Altura", true
		}
	}
	return "none", false
}
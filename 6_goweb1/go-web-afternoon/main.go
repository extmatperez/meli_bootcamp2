package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


type Empleado struct {
	Id int  `json:"id"`
	Name string  `json:"name"`
	Active string  `json:"active"`
}

var empleados = []Empleado{
	{1, "Juan Cruz", "active"},
	{2, "Donald", "active"},
	{3, "Boris", "inactive"},
}

func filteringByActive(c *gin.Context) {

	filter := c.Query("active")

	var filtrados []Empleado

	for _, v := range empleados {
		if filter == v.Active {
			filtrados = append(filtrados, v)
		}
	}

	if len(filtrados) > 0 {
		c.JSON(http.StatusOK, filtrados)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	}

}

func filteringById(c *gin.Context) {
	filter := c.Param("id")
	
	var empleado []*Empleado

	for _, v := range empleados {
		filtered_id, _ := strconv.Atoi(filter)
		if filtered_id == v.Id {
			empleado = append(empleado, &v)
			break
		}
	}

	if len(empleado) > 0 {
		c.JSON(http.StatusOK, empleado)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	}

}

func main() {

	r := gin.Default()
	empleadosEP := r.Group("/empleados")
	{
		empleadosEP.GET("/", filteringByActive)
		empleadosEP.GET("/:id", filteringById)
	}
	r.Run()

}
package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
}

func AddPerson(c *gin.Context) {
	var per Person
	err := c.ShouldBindJSON(&per)
	if err != nil {
		// c.JSON(400,gin.H{
		// 	"error":err.Error(),
		// })
		c.String(400, "An error ocurred: %v", err.Error())
	} else {
		per.ID = 1 //si existe lo pisa
		c.JSON(200, per)
	}
}
func main() {
	router := gin.Default()
	router.POST("/person/add", AddPerson)
	router.Run()
}

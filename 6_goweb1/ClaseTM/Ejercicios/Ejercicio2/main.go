package main

import "github.com/gin-gonic/gin"

type Fecha struct {
	Dia  int
	Mes  int
	Anio int
}
type Producto struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion Fecha   `json:"fecha_creacion"`
}

func handlerSaludar(c *gin.Context) {
	saludandoA := "Hola " + c.Param("nombre")

	c.JSON(200, gin.H{
		"message": saludandoA,
	})
}
func main() {
	router := gin.Default()

	router.GET("api/Hello/:nombre", handlerSaludar)

	// http://localhost:8080/api/Hello/Nahuel
	// {"message":"Hola Nahuel"}

	router.Run()
}

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	Publicado     bool    `json:"publicado" binding:"required"`
	FechaCreacion Fecha   `json:"fecha_creacion" binding:"required"`
}

var ArrProductos []Producto
var Token string = "HYUOWIE17235312"

func CargarData(c *gin.Context) {
	data, err := os.ReadFile("/Users/nscerca/Desktop/meli_bootcamp/meli_bootcamp2/6_goweb1/Productos.json")
	if err == nil {
		json.Unmarshal(data, &ArrProductos)

		c.String(200, "Productos Cargados")

	} else {
		c.JSON(4040, gin.H{
			"message": "No se cargaron los datos.",
		})
	}
}

func getWithId(c *gin.Context) {
	idToSearch := c.Param("id")

	var prodFiltrados []Producto

	for i, item := range ArrProductos {
		if idToSearch == strconv.Itoa(item.ID) {
			prodFiltrados = append(prodFiltrados, ArrProductos[i])
		}
	}

	if len(prodFiltrados) != 0 {
		c.JSON(200, prodFiltrados)

	} else {
		c.JSON(4040, gin.H{
			"message": "No se encontraron los datos solicitados.",
		})
	}

}

func getAll(c *gin.Context) {
	token := c.GetHeader("token")
	if token == Token {

		if len(ArrProductos) != 0 {
			c.JSON(200, ArrProductos)
		} else {
			c.JSON(404, gin.H{
				"message": "No hay datos.",
			})
		}

	} else {
		c.JSON(401, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
	}

}
func validar(req Producto) string {
	r := reflect.ValueOf(req)
	for i := 0; 1 < r.NumField(); i++ {
		varValor := r.Field(i).Interface()
		s := reflect.TypeOf(varValor).Kind()

		fmt.Println(s)

		// if fmt.Sprintf("%s",s)  {

		// }
	}
	return ""
}
func postProducto(c *gin.Context) {

	var prod Producto
	err := c.ShouldBindJSON(&prod)
	token := c.GetHeader("token")

	// req := validar(prod)
	// fmt.Println(req)

	// jsonData, er := c.GetRawData()
	// if er == nil {
	// 	fmt.Println(jsonData)

	// }

	if token == Token {
		if err == nil {
			if len(ArrProductos) == 0 {
				prod.ID = 1
			} else {

				idInt := int(ArrProductos[len(ArrProductos)-1].ID)
				prod.ID = idInt + 1
			}

			ArrProductos = append(ArrProductos, prod)
			c.JSON(201, prod)

		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		c.JSON(401, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
	}

}

func getToken(c *gin.Context) {
	type User struct {
		Mail string
		Pass string
	}
	var us User
	err := c.ShouldBindJSON(&us)

	if err == nil {
		if us.Mail == "nahuel@gmail.com" && us.Pass == "1234" {
			c.JSON(200, gin.H{
				"token": Token,
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Fallo el Logeo",
			})
		}

	} else {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

}

func main() {
	router := gin.Default()
	groupProducts := router.Group("api/productos")
	{
		groupProducts.POST("/add", postProducto)
		groupProducts.GET("/", getAll)
		groupProducts.GET(":id", getWithId)
	}

	router.POST("/api/getToken", getToken)
	router.GET("/api/loadData", CargarData)
	router.Run()
}

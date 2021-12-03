package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre" binding:"required"`
	Color         string `json:"color" binding:"required"`
	Precio        int    `json:"precio" binding:"required"`
	Stock         int    `json:"stock" binding:"required"`
	Codigo        string `json:"codigo" binding:"required"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion" binding:"required"`
}

var products []Product

func LoadData() {
	content, err := os.ReadFile("products.json")

	if err != nil {
		fmt.Println(err)
		panic("el archivo indicado no fue encontrado o está dañado\n")
	}

	p := []Product{}

	json.Unmarshal(content, &p)

	products = p
}
func GetAll(c *gin.Context) {
	c.JSON(200, products)
}

func Filter(c *gin.Context) {

	var filtrados []*Product

	for i := 0; i < len(products); i++ {
		var coincide []bool

		/*if c.Query("id") == "" || c.Query("id") == products[i].ID {
			coincide = append(coincide, true)
		}*/
		if c.Query("nombre") == "" || c.Query("nombre") == products[i].Nombre {
			coincide = append(coincide, true)
		}

		if c.Query("color") == "" || c.Query("color") == products[i].Color {
			coincide = append(coincide, true)
		}

		if c.Query("precio") == "" {
			coincide = append(coincide, true)
		} else {
			filtroPrecio, err := strconv.Atoi(c.Query("precio"))
			if err == nil && filtroPrecio == products[i].Precio {
				coincide = append(coincide, true)
			}
		}
		if c.Query("stock") == "" {
			coincide = append(coincide, true)
		} else {
			filtroStock, err := strconv.Atoi(c.Query("stock"))
			if err == nil && filtroStock == products[i].Stock {
				coincide = append(coincide, true)
			}
		}

		if c.Query("codigo") == "" || c.Query("codigo") == products[i].Codigo {
			coincide = append(coincide, true)
		}
		if c.Query("publicado") != "" {
			if c.Query("publicado") == "true" && products[i].Publicado {
				coincide = append(coincide, true)
			} else if c.Query("publicado") == "false" && !products[i].Publicado {
				coincide = append(coincide, true)
			}
		} else if c.Query("publicado") == "" {
			coincide = append(coincide, true)
		}
		if c.Query("fechaCreacion") == "" || c.Query("fechaCreacion") == products[i].FechaCreacion {
			coincide = append(coincide, true)
		}
		//	fmt.Printf("len= %d idCoincide=", len(coincide), c.Query("id") == products[i].ID)
		if len(coincide) > 7 {
			filtrados = append(filtrados, &products[i])
		}
	}

	c.JSON(200, filtrados)

}

func FindById(c *gin.Context) {

	finded := false
	id, err := strconv.Atoi(c.Param("id"))
	for _, value := range products {

		if err == nil && value.ID == id {
			c.JSON(200, value)
			finded = true
			break
		}
	}
	if !finded {
		c.String(404, "Producto no encontrado")
	}
	mapQuery := c.Request.URL.Query()

	for i, v := range mapQuery {
		fmt.Println("Key -> ", i, " value -> ", v[0])
	}

}

func Register(c *gin.Context) {

	token := c.GetHeader("token")
	if token != "1234" {
		c.String(http.StatusUnauthorized, "no tiene permisos para realizar la petición solicitada")
		return
	}

	var newProd Product
	err := c.ShouldBindJSON(&newProd)

	if err != nil {
		if strings.Contains(err.Error(), "required") {
			tipos := reflect.TypeOf(newProd)
			i := 0
			var errores []string
			for i = 0; i < tipos.NumField(); i++ {
				if strings.Contains(err.Error(), tipos.Field(i).Name) {
					errores = append(errores, fmt.Sprintf("Error: el campo %s es requerido", tipos.Field(i).Name))
				}
			}
			if len(errores) == 1 {
				c.JSON(400, gin.H{
					"error": errores[0],
				})
			} else {
				c.JSON(400, errores)
			}
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}

		return
	}

	if len(products) == 0 {
		newProd.ID = 1
	} else {
		newProd.ID = products[len(products)-1].ID + 1
	}
	products = append(products, newProd)
	c.JSON(http.StatusOK, newProd)

}

func filterByParams(paramas map[string][]string) {

	for key, element := range paramas {

	}
}

func filterByParam(param string, element []string, p []Product) []Product {
	prod := Product{}
	tipos := reflect.TypeOf(prod)
	fieldNum := -1
	for i := 0; i < tipos.NumField(); i++ {
		if strings.ToLower(param) == strings.ToLower(tipos.Field(i).Name) {
			fieldNum = i
			break
		}
	}

	for i, v := range p {
		
		if 
	}



}

func main() {
	LoadData()
	router := gin.Default()
	productsRoute := router.Group("products")
	productsRoute.GET("", GetAll)
	productsRoute.GET("/filter", Filter)
	productsRoute.GET("/:id", FindById)
	productsRoute.POST("", Register)
	router.Run()
}

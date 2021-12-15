package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Ejercicio 1 - Test Unitario GetAll()
Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
1. Dentro de la carpeta /internal/products, crear un archivo repository_test.go con el test diseñado.
Ejercicio 2 - Test Unitario UpdateName()
Diseñar Test de UpdateName, donde se valide que la respuesta retornada sea correcta para la actualización del nombre de un producto específico. Y además se compruebe que efectivamente se usa el método “Read” del Storage para buscar el producto. Para esto:
1. Crear un mock de Storage, dicho mock debe contener en su data un producto específico cuyo nombre puede ser “Before Update”.
2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. Puede ser a través de un boolean como se observó en la clase.
3. Para dar el test como OK debe validarse que al invocar el método del Repository UpdateName, con el id del producto mockeado y con el nuevo nombre “After Update”, efectivamente haga la actualización. También debe validarse que el método Read haya sido ejecutado durante el test.
*/

type Product struct {
	ID              int     `json:"id"`
	Nombre          string  `json:"nombre"`
	Color           string  `json:"color"`
	Precio          float64 `json:"precio"`
	Stock           int     `json:"stock"`
	Codigo          string  `json:"codigo"`
	Publicado       bool    `json:"publicado"`
	FechaDeCreacion string  `json:"fecha_de_creacion"`
}

func Ejemplo(context *gin.Context) {
	header := context.Request.Header
	fmt.Println("Header", header)
}
func Name(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
func GetID(c *gin.Context) {
	id := c.Param("id")
	var products []Product
	jsonFile, err := os.ReadFile("./products.json")
	if err == nil {
		json.Unmarshal(jsonFile, &products)
		idInt, _ := strconv.Atoi(id)
		flag := false
		for _, product := range products {
			if product.ID == idInt {
				flag = true
				c.JSON(http.StatusOK, product)
			}
		}
		if flag == false {
			c.String(404, "404 producto no encontrado")
		}
	} else {
		c.String(200, "404 producto no encontrado")
	}
}

func GetAll(c *gin.Context) {
	var products []interface{}
	jsonFile, err := os.ReadFile("./products.json")
	if err == nil {
		json.Unmarshal(jsonFile, &products)
		c.JSON(http.StatusOK, products)
		// c.JSON(http.StatusOK, gin.H{
		// 	"products": listProducts,
		// })
	}
}

func main() {
	router := gin.Default()
	router.GET("/products/:name", Name)
	router.GET("/producto/:id", GetID)
	router.GET("/products/All", GetAll)
	router.Run(":8080")
}

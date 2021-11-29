package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID            string `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

func LoadData() []Product {
	content, err := os.ReadFile("products.json")

	if err != nil {
		fmt.Println(err)
		panic("el archivo indicado no fue encontrado o está dañado\n")
	}

	p := []Product{}

	json.Unmarshal(content, &p)

	return p
}
func GetAll(c *gin.Context) {

	p := LoadData()
	c.JSON(200, p)
}

func Filter(c *gin.Context) {

	products := LoadData()

	var filtrados []*Product

	for i := 0; i < len(products); i++ {
		var coincide []bool

		if c.Query("id") == "" || c.Query("id") == products[i].ID {
			coincide = append(coincide, true)
		}
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

	products := LoadData()
	finded := false
	for _, value := range products {
		if value.ID == c.Param("id") {
			c.JSON(200, value)
			finded = true
			break
		}
	}
	if !finded {
		c.String(404, "Producto no encontrado")
	}


}

func main() {
	router := gin.Default()

	router.GET("/products", GetAll)
	router.GET("/products/filter", Filter)
	router.GET("/products/:id", FindById)
	router.Run()
}

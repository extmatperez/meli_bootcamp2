package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreacion string  `json:"fechaCreacion"`
}

var productosListo []Producto

func saludo(c *gin.Context) {
	nombre := c.Param("nombre")
	c.JSON(200, gin.H{
		"mensaje": "Hola, " + nombre,
	})
}
func GetAll(c *gin.Context) {

	dbproductos, _ := ioutil.ReadFile("products.json")
	err := json.Unmarshal(dbproductos, &productosListo)

	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, productosListo)
		//c.String(200, productosListo)	//otra forma de devolverlo
	}
}
func Ejemplo(ctx *gin.Context) {
	contenido := ctx.Request.Body
	header := ctx.Request.Header
	metodo := ctx.Request.Method

	fmt.Println("Recibi algo")
	fmt.Println("Metodo ", metodo)
	fmt.Println("Cabecera ")

	for k, v := range header {
		fmt.Println(k, ":", v)
	}
	fmt.Println("Contenido ", contenido)
	//ctx.JSON(200, "Salida")
	ctx.String(200, "Termine")
}
func filtraId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var filtrados []Producto

	dbproductos, _ := ioutil.ReadFile("products.json")
	err := json.Unmarshal(dbproductos, &productosListo)

	if err != nil {
		c.String(400, "Algo salió mal")
	} else {
		for _, e := range productosListo {
			if id == e.Id {
				filtrados = append(filtrados, e)
			}
		}
		if len(filtrados) > 0 {
			c.JSON(200, filtrados)
		} else {
			c.String(404, "No se encontro el producto")
		}
	}

	/* if c.BindQuery(&productos1) == nil {
		parametro := c.Param("id")
		for _, v := range productos1 {
			if v.Id == parametro {
				c.JSON(200, v)
			}
		}
	} */

}
func filtraNombre(c *gin.Context) {
	nombre := c.Query("nombre")
	var productosListo []Producto
	var filtrados []Producto

	dbproductos, _ := ioutil.ReadFile("./products.json")
	err := json.Unmarshal(dbproductos, &productosListo)

	if err != nil {
		c.String(400, "Algo salió mal")
	} else {
		for _, prod := range productosListo {
			if strings.Contains(prod.Nombre, nombre) {
				filtrados = append(filtrados, prod)
			}
		}
		if len(filtrados) > 0 {
			c.JSON(200, filtrados)
		} else {
			c.String(404, "No se encontraron coincidencias")
		}
	}
}
func filtraPrecio(c *gin.Context) {
	max, _ := strconv.ParseFloat(c.Query("max"), 64)
	min, _ := strconv.ParseFloat(c.Query("min"), 64)
	var productosListo []Producto
	var filtrados []Producto

	dbproductos, _ := ioutil.ReadFile("./products.json")
	err := json.Unmarshal(dbproductos, &productosListo)

	if err != nil {
		c.String(400, "Algo salió mal")
	} else {
		for _, prod := range productosListo {
			if max >= prod.Precio && min <= prod.Precio {
				filtrados = append(filtrados, prod)
			}
		}
		if len(filtrados) > 0 {
			c.JSON(200, filtrados)
		} else {
			c.String(404, "No se encontraron productos en ese rango")
		}
	}
}
func AddPersona(c *gin.Context) {

	var prod Producto
	err := c.ShouldBind(&prod)
	token := c.GetHeader("token")

	if err != nil {
		c.String(400, "Algo salió mal")
		return
	}

	if token != "123" {
		c.String(401, "no tiene permisos para realizar la petición solicitada")
		return
	} else {

		var falla string
		switch {
		case prod.Nombre == "":
			falla = "Nombre"
		case prod.Color == "":
			falla = "Color"
		case prod.Precio == 0:
			falla = "Precio"
		case prod.Stock == 0:
			falla = "Stock"
		case prod.Codigo == "":
			falla = "Codigo"
		case prod.Publicado == false:
			falla = "Publicado"
		case prod.FechaCreacion == "":
			falla = "FechaCreacion"
		default:
			falla = "todoBien"
		}

		if falla == "todoBien" {

			ultimo := productosListo[len(productosListo)-1].Id + 1
			prod.Id = ultimo
			productosListo = append(productosListo, prod)
			c.JSON(200, prod)
		} else {

		}
	}

	/* tipos := reflect.TypeOf(prod)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		// fmt.Println(i, "->", tipos.Field(i).Name)
		if strings.ToLower(tipos.Field(i).Name) == campo {
			break
		}
	}

	return filtrado */

}
func GetActual(c *gin.Context) {

	c.JSON(200, productosListo)
	//c.String(200, productosListo)	//otra forma de devolverlo

}
func main() {

	router := gin.Default()

	router.GET("/hola/:nombre", saludo)
	router.GET("/productos", GetAll)
	router.GET("/productosActual", GetActual)
	router.GET("/productos/:id", filtraId)

	router.GET("/ejemplo", Ejemplo)

	grupoFiltrador := router.Group("/filtrar")
	{
		grupoFiltrador.GET("/nombre", filtraNombre)
		grupoFiltrador.GET("/precios", filtraPrecio)
	}

	router.POST("/productos", AddPersona)
	router.Run()
}

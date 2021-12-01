package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        int    `json:"precio"`
	Stock         string `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fecha_creacion"`
}

var products []Products

func loadData(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "" {
		if token == "123" {
			ps, err := os.ReadFile("./7_goweb2/go-web/archivos/products.json")
			if err != nil {
				fmt.Println(err)
			} else {
				json.Unmarshal(ps, &products)
			}
			c.JSON(http.StatusOK, gin.H{"success": "productos cargados"})
		} else {
			c.JSON(401, gin.H{"error": "no tiene permiso para realziar la peticion solicitada"})
		}
	} else {
		c.JSON(401, gin.H{"error": "es necesario autenticarse para continuar"})
	}
}
func showData(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}
func add(c *gin.Context) {
	var p Products
	err := c.ShouldBindJSON(&p) // se usa cuando viene del body
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		if len(products) == 0 {
			p.ID = 1
		} else {
			p.ID = products[len(products)-1].ID + 1
		}
		products = append(products, p)
		c.JSON(200, p)
	}
}

func filtrar(sliceProductos []Products, campo string, valor string) []Products {
	var filtrado []Products
	var p Products
	tipos := reflect.TypeOf(p)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		// fmt.Println(i, "->", tipos.Field(i).Name)
		if strings.ToLower(tipos.Field(i).Name) == campo {
			break
		}
	}
	for _, v := range sliceProductos {
		//var cadena string
		cadena := fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor) {
			// if reflect.ValueOf(v).Field(i).Interface() == valor {
			filtrado = append(filtrado, v)
		}
	}
	return filtrado
}
func validar(p Products) string {
	fmt.Println(p)
	return ""
}

func FiltrarProductos(ctx *gin.Context) {
	var etiquetas []string
	var p Products
	var val string
	var productosFiltrados []Products
	productosFiltrados = products
	err := ctx.ShouldBind(&p) // este se usa cuando viene de form
	etiquetas = append(etiquetas, "nombre", "color", "stock", "codigo", "fecha_creacion", "precio", "publicado")

	if err != nil {
		val = "Error: con el formato de campos invalido"
		fmt.Println(err)

	} else {
		val = validar(p)
		fmt.Println(p)
	}
	if val != "" {
		for _, v := range etiquetas {
			// fmt.Println(v, "->", ctx.Query(v))
			//if ctx.Query(v) == "" {
			//	salida += fmt.Sprintf("el campo %s no puede estar vacio \n", v)
			//}
			if len(ctx.Query(v)) != 0 && len(productosFiltrados) != 0 {
				productosFiltrados = filtrar(productosFiltrados, v, ctx.Query(v))
			}
		}
		//ctx.String(200, val)

		if len(productosFiltrados) == 0 {
			ctx.String(200, "No hay coincidencias")
		} else {
			ctx.JSON(200, productosFiltrados)
		}
	}
}

func main() {

	router := gin.Default()
	products := router.Group("/products")
	products.GET("/", loadData)
	products.GET("/show", showData)
	products.POST("/add", add)
	products.GET("/filtros", FiltrarProductos)
	//router.GET("/products", GetAll)
	//router.GET("product/:id", GetById)             //localhost:8080/product/1
	//router.GET("productpublicado", GetByPublicado) // localhost:8080/productpublicado?publicado=false

	router.Run()
}

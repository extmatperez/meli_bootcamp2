package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Color        string `json:"color"`
	Price        string `json:"price"`
	Stock        string `json:"stock"`
	Code         string `json:"code"`
	Published    bool   `json:"published"`
	CreationDate string `json:"creation_date"`
}

func readData() []Products {

	var list []Products
	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal([]byte(readProducts), &list); err != nil {
		log.Fatal(err)
	}
	return list
}

func getAll(c *gin.Context) {

	var prodList = readData()

	c.JSON(200, gin.H{
		"data": prodList,
	})
}

func getOne(c *gin.Context) {
	parameter := c.Param("id")

	var prodList = readData()

	//var prod []Products
	prods := []Products{}
	//is_Product := false

	for _, v := range prodList {
		if strconv.Itoa(v.ID) == parameter {
			prods = append(prods, v)
			//is_Product = true
		}
	}

	if len(prods) > 0 {
		c.JSON(200, prods)
	} else {
		c.String(400, "No product found")
	}
}

func filterProducts(ctx *gin.Context) {
	var filtered []*Products
	prodList := readData()

	for i, v := range prodList {
		if ctx.Query("filter") == strconv.FormatBool(v.Published) {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Name {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Color {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Price {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Stock {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.Code {
			filtered = append(filtered, &prodList[i])
		} else if ctx.Query("filter") == v.CreationDate {
			filtered = append(filtered, &prodList[i])
		}
	}

	if len(filtered) != 0 {
		ctx.JSON(200, filtered)
	} else {
		ctx.String(400, "No results found")
	}

}
func validation(req Products) string {
	reqValue := reflect.ValueOf(req)
	for i := 0; i < reqValue.NumField(); i++ {
		value := reqValue.Field(i).Interface()
		tipe := reflect.TypeOf(value).Kind()

		if fmt.Sprint(tipe) == "string" {
			if value == "" {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		} else if fmt.Sprint(tipe) == "int64" {
			if value.(int64) == 0 {
				return fmt.Sprintf("El campo %v no puede estar vacío", reqValue.Type().Field(i).Name)
			}
		}

	}
	return ""
}

func addProduct(ctx *gin.Context) {
	var prod Products
	prodList := readData()
	err := ctx.ShouldBindJSON(&prod)
	lenthProds := len(prodList)
	//tipes := reflect.TypeOf(prodList)
	prodValidate := validation(prod)
	token := ctx.GetHeader("token")
	if prodValidate != "" {
		ctx.String(400, prodValidate)
		return
	}

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	} else {
		if lenthProds == 0 {
			prod.ID = 1
		} else {
			prod.ID = prodList[lenthProds-1].ID + 1
		}
		if token != "" {
			if token == "123456" {
				if err != nil {
					ctx.String(400, "Ha ocurrido un error")
				}
				prodList = append(prodList, prod)
				ctx.JSON(200, prod)
			} else {
				ctx.String(401, "No tiene permisos para realizar la peticion solicitada")
			}
		} else {
			ctx.String(400, "No ingreso token")
		}
	}
}

func main() {
	// Crea un router con gin
	router := gin.Default()
	// Captura la solicitud GET “/hello-world”
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Franco!",
		})
	})
	//router.GET("/find/products", filterProducts)

	products := router.Group("/products")
	products.GET("/find", filterProducts)
	products.GET("/:id", getOne)
	products.GET("/", getAll)
	products.POST("/", addProduct)
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()

}

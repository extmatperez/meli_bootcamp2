package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Products struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Color     string  `json:"color"`
	Price     float64 `json:"price"`
	Stock     int64   `json:"stock"`
	Code      int64   `json:"code"`
	Published string  `json:"published"`
	Created   string  `json:"created"`
}

func readData() []Products {

	var list []Products
	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal([]byte(readProducts), &list); err != nil {
		log.Fatal(err)
	}
	return list
}

///// FUNCIONES HANDLERS ///////

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Estefi!",
	})
}

func getAll(c *gin.Context) {

	var prodList = readData()

	c.JSON(200, gin.H{
		"data": prodList,
	})
}
func getbyFilter(c *gin.Context) {

	var prodList = readData()
	var filtrados []Products

	selectedFilter := c.Request.URL.Query().Get("filterValue")
	switch c.Request.URL.Query().Get("filter") {
	case "name":
		for _, p := range prodList {
			if selectedFilter == p.Name {
				filtrados = append(filtrados, p)
			}
		}

		c.JSON(200, gin.H{
			"data": filtrados,
		})
	case "color":
		for _, p := range prodList {
			if selectedFilter == p.Color {
				filtrados = append(filtrados, p)
			}
		}
		c.JSON(200, gin.H{
			"data": filtrados,
		})
	case "price":
		i, _ := strconv.ParseFloat(c.Query(selectedFilter), 64)

		for _, p := range prodList {
			if i == p.Price {
				filtrados = append(filtrados, p)
			}
		}
		c.JSON(200, gin.H{
			"data": filtrados,
		})
	case "stock":
		i, _ := strconv.ParseInt(c.Query(selectedFilter), 10, 64)

		for _, p := range prodList {
			if i == p.Stock {
				filtrados = append(filtrados, p)
			}
		}
		c.JSON(200, gin.H{
			"data": filtrados,
		})
	case "code":
		i, _ := strconv.ParseInt(c.Query(selectedFilter), 10, 64)

		for _, p := range prodList {
			if i == p.Code {
				filtrados = append(filtrados, p)
			}
		}
		c.JSON(200, gin.H{
			"data": filtrados,
		})
	case "published":
		for _, p := range prodList {
			if selectedFilter == p.Published {
				filtrados = append(filtrados, p)
			}
		}
		c.JSON(200, gin.H{
			"data": filtrados,
		})

	}

}

// func getbyName(c *gin.Context) {

// 	var prodList = readData()
// 	var filtrados []Products

// 	for _, p := range prodList {
// 		if c.Query("name") == p.Name {
// 			filtrados = append(filtrados, p)
// 		}
// 	}

// 	c.JSON(200, gin.H{
// 		"data": filtrados,
// 	})

// }

// func getbyColor(c *gin.Context) {

// 	var prodList = readData()
// 	var filtrados []Products

// 	for _, p := range prodList {
// 		if c.Query("color") == p.Color {
// 			filtrados = append(filtrados, p)
// 		}
// 	}

// 	c.JSON(200, gin.H{
// 		"data": filtrados,
// 	})

// }

// func getbyPrice(c *gin.Context) {

// 	var prodList = readData()
// 	var filtrados []Products

// 	for _, p := range prodList {
// 		i, _ := strconv.ParseFloat(c.Query("price"), 64)
// 		if i == p.Price {
// 			filtrados = append(filtrados, p)
// 		}
// 	}

// 	c.JSON(200, gin.H{
// 		"data": filtrados,
// 	})

// }

func getbyId(c *gin.Context) {

	var prodList = readData()
	var selectedProd []Products

	prodId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	for _, p := range prodList {
		if prodId == p.ID {
			selectedProd = append(selectedProd, p)
		}
	}
	c.JSON(200, gin.H{
		"data": selectedProd,
	})
}

func main() {

	router := gin.Default()

	router.GET("/", sayHello)
	router.GET("/products", getAll)
	router.GET("/products/:id", getbyId)
	router.GET("/products/filter/select", getbyFilter)

	// productfilter := router.Group("/productfilter")
	// {
	// 	productfilter.GET("/name", getbyName)
	// 	productfilter.GET("/color", getbyColor)
	// 	productfilter.GET("/price", getbyPrice)
	// 	productfilter.GET("/stock", getbyStock)
	// 	productfilter.GET("/code", getbyCode)
	// 	productfilter.GET("/published", getbyPublished)
	// }

	router.Run()
}

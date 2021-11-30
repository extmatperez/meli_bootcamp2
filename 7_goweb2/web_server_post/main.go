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

var prodList []Products

var lastID int64

var tokenPrueba string

func readData() {

	readProducts, _ := os.ReadFile("./products.json")

	if err := json.Unmarshal(readProducts, &prodList); err != nil {
		log.Fatal(err)
	}

}

///// HANDLERS ///////

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home",
	})
}

///// POST HANDLERS ///////
func addProduct(c *gin.Context) {
	var req Products

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	lastID++
	req.ID = lastID
	prodList = append(prodList, req)
	c.JSON(200, req)
}

///// GET HANDLERS ///////

func getAll(c *gin.Context) {

	// page, _ := strconv.ParseInt(c.Request.URL.Query().Get("page"), 10, 64)
	// limit, _ := strconv.ParseInt(c.Request.URL.Query().Get("limit"), 10, 64)

	// startIndex := (page - 1) * limit
	// endIndex := page * limit

	// var paginatedResults []Products
	// paginatedResults = prodList[startIndex:endIndex]
	token := c.GetHeader("token")
	if token != tokenPrueba {
		c.JSON(401, gin.H{
			"error": "token inválido",
		})
	} else {

		c.JSON(200, gin.H{
			"data": prodList,
		})
	}

}

// func getbyFilter(c *gin.Context) {

// 	var filtrados []Products

// 	selectedFilter := c.Request.URL.Query().Get("filterValue")
// 	switch c.Request.URL.Query().Get("filter") {
// 	case "name":
// 		for _, p := range prodList {
// 			if selectedFilter == p.Name {
// 				filtrados = append(filtrados, p)
// 			}
// 		}

// 		c.JSON(200, gin.H{
// 			"data": filtrados,
// 		})
// 	case "color":
// 		for _, p := range prodList {
// 			if selectedFilter == p.Color {
// 				filtrados = append(filtrados, p)
// 			}
// 		}
// 		c.JSON(200, gin.H{
// 			"data": filtrados,
// 		})
// 	case "price":
// 		i, _ := strconv.ParseFloat(c.Query(selectedFilter), 64)

// 		for _, p := range prodList {
// 			if i == p.Price {
// 				filtrados = append(filtrados, p)
// 			}
// 		}
// 		c.JSON(200, gin.H{
// 			"data": filtrados,
// 		})
// 	case "stock":
// 		i, _ := strconv.ParseInt(c.Query(selectedFilter), 10, 64)

// 		for _, p := range prodList {
// 			if i == p.Stock {
// 				filtrados = append(filtrados, p)
// 			}
// 		}
// 		c.JSON(200, gin.H{
// 			"data": filtrados,
// 		})
// 	case "code":
// 		i, _ := strconv.ParseInt(c.Query(selectedFilter), 10, 64)

// 		for _, p := range prodList {
// 			if i == p.Code {
// 				filtrados = append(filtrados, p)
// 			}
// 		}
// 		c.JSON(200, gin.H{
// 			"data": filtrados,
// 		})
// 	case "published":
// 		for _, p := range prodList {
// 			if selectedFilter == p.Published {
// 				filtrados = append(filtrados, p)
// 			}
// 		}
// 		c.JSON(200, gin.H{
// 			"data": filtrados,
// 		})

// 	}

// }

func getbyId(c *gin.Context) {

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

	readData()
	length := len(prodList) - 1
	lastID = prodList[length].ID
	tokenPrueba = "1234"

	/////////////////////// start server /////////////////////////////
	router := gin.Default()

	router.GET("/", sayHello)

	products := router.Group("/products")

	products.GET("/", getAll)
	products.GET("/products/:id", getbyId)
	// products.GET("/products/filter/select", getbyFilter)

	products.POST("/addproduct", addProduct)

	router.Run()

}

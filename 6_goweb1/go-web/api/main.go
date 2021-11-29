package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/sayHi/:name/:lastName", func(c *gin.Context) {
		name := c.Param("name")
		lastName := c.Param("lastName")
		greeting := fmt.Sprintf("Hi, %v %v", name, lastName)
		c.JSON(http.StatusOK, greeting)
	})

	router.GET("/transactions/GetAll", func(c *gin.Context) {
		jsonFile, _ := os.Open("transactions.json")
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var transactions []Transaction
		err := json.Unmarshal(byteValue, &transactions)
		if err == nil {
			c.JSON(http.StatusOK, transactions)
		} else {
			c.JSON(http.StatusOK, transactions)
		}
	})

	router.Run()
}

type Transaction struct {
	ID              int     `json:"id"`
	TransactionCode string  `json:"transactionCode"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Sender          string  `json:"sender"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transactionDate"`
}

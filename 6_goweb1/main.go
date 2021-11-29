package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Greeting struct {
	Message string `json:"message"`
}

func (g *Greeting) setMessage(name string) {
	g.Message = fmt.Sprintf("Hola %v", name)
}

type Transaction struct {
	Id              int    `json:"id"`
	TransactionCode string `json:"transactionCode"`
	Currency        string `json:"currency"`
	Amount          int    `json:"amount"`
	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	Date            string `json:"date"`
}

func main() {
	router := gin.Default()

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")

		newGreeting := Greeting{}

		newGreeting.setMessage(name)

		c.JSON(http.StatusOK, newGreeting)
	})

	router.GET("/transactions/getAll", func(c *gin.Context) {

		data, err := os.ReadFile("./transactions.json")

		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		var transactions = []Transaction{}

		unmarshalErr := json.Unmarshal(data, &transactions)
		if unmarshalErr != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, transactions)
	})

	router.Run()
}

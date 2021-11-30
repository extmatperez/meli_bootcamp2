package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	transacionFile := "../../transaction.json"
	data, err := os.ReadFile(transacionFile)

	if err != nil {
		fmt.Printf("There was a error %v", err)
	}

	transactions := toDeserializer(data)

	//fmt.Printf("%+v\n", transactions)

	router := gin.Default()

	router.GET("/transactions", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"transaction": transactions,
		})
	})

	router.Run()
}

type transaction struct {
	ID              int     `json:"id"`
	TransactionCode string  `json:"transaction_code"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Receiver        string  `json:"receiver"`
	Sender          string  `json:"sender"`
	TransactionDate string  `json:"transaction_date"`
}

func toDeserializer(data []byte) []transaction {
	var transactions []transaction

	if err := json.Unmarshal(data, &transactions); err != nil {
		panic(err)
	}

	return transactions
}

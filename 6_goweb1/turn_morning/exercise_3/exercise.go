package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//fmt.Printf("%+v\n", transactions)

	router := gin.Default()

	router.GET("/transactions/:id", getAll)

	router.Run()

	/* http.HandleFunc("/transactions", getAll)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Listening in 8080...") */
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

func readTransactions() []transaction {
	transacionFile := "../../transaction.json"
	data, err := os.ReadFile(transacionFile)

	if err != nil {
		fmt.Printf("There was a error %v", err)
	}

	return toDeserializer(data)
}

func toDeserializer(data []byte) []transaction {
	var transactions []transaction

	if err := json.Unmarshal(data, &transactions); err != nil {
		panic(err)
	}

	return transactions
}

func getAll(c *gin.Context) {

	transactions := readTransactions()
	c.JSON(200, gin.H{
		"transaction": transactions})
}

func getAll3(w http.ResponseWriter, r *http.Request) {

	transactions := readTransactions()

	for _, v := range transactions {
		fmt.Fprintf(w, v.Currency)

	}
}

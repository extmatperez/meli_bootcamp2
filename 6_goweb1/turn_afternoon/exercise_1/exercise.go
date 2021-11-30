package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//fmt.Printf("%+v\n", transactions)

	router := gin.Default()

	router.GET("/transactions", getAll)
	router.GET("/transactions/receivers/:receiver", getByReceiver)

	router.Run(":9090")

	/* http.HandleFunc("/transactions", getAll)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Listening in 8080...") */
}

type transaction struct {
	ID              int     `form:"id" json:"id"`
	TransactionCode string  `form:"transaction_code" json:"transaction_code"`
	Currency        string  `form:"currency" json:"currency"`
	Amount          float64 `form:"amount" json:"amount"`
	Receiver        string  `form:"receiver" json:"receiver"`
	Sender          string  `form:"sender" json:"sender"`
	TransactionDate string  `form:"transaction_date" json:"transaction_date"`
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

func getAll(context *gin.Context) {

	transactions := readTransactions()
	context.JSON(200, gin.H{
		"transaction": transactions})
}

func getByReceiver(context *gin.Context) {

	var idParam string = context.Param("receiver")
	transactions := readTransactions()
	var transFound []*transaction
	for _, trans := range transactions {
		if idParam == trans.Receiver {
			transFound = append(transFound, &trans)
			break
		}
	}
	if transFound != nil {
		context.JSON(200, gin.H{
			"transaction": transFound})
	} else {
		context.String(404, "Receiver %s not found", idParam)

	}

}

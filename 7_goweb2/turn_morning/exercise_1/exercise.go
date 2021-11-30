package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var transactions []transaction
var globalID int

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

	context.JSON(200, gin.H{
		"transaction": transactions})
}

func getById(context *gin.Context) {

	var gotTrans []*transaction
	idParam, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		panic(err)
	}
	for _, trans := range transactions {
		if idParam == trans.ID {
			gotTrans = append(gotTrans, &trans)
			break
		}
	}
	if gotTrans != nil {
		context.JSON(200, gin.H{
			"transaction": gotTrans})
	} else {
		context.String(404, "Transaction %v not found", idParam)
	}

}

func getByReceiver(context *gin.Context) {

	var idParam string = context.Param("receiver")
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

func createTransaction(context *gin.Context) {

	var newTransactions transaction
	err := context.ShouldBindJSON(&newTransactions)
	if err != nil {
		context.String(200, "{ \n \"error\": \"%v\", err")
	} else {
		if len(transactions) == 0 {
			newTransactions.ID = 1
		} else {
			newTransactions.ID = transactions[len(transactions)-1].ID + 1
		}
		transactions = append(transactions, newTransactions)

		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(transactions)

		data := reqBodyBytes.Bytes()

		//data := []byte(fmt.Sprintf("%v", transactions))
		err := os.WriteFile("../../transaction.json", data, 0644)
		if err != nil {
			context.JSON(400, gin.H{"error": err})
		} else {
			context.JSON(200, gin.H{"transaction": newTransactions})
		}
	}

}

func main() {

	//fmt.Printf("%+v\n", transactions)

	router := gin.Default()

	transactions = readTransactions()

	transactionURL := router.Group("/transactions")

	transactionURL.GET("/", getAll)
	transactionURL.GET("/:id", getById)
	transactionURL.GET("/receivers/:receiver", getByReceiver)
	transactionURL.POST("/", createTransaction)

	router.Run(":9090")

}

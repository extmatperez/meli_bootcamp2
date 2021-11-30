package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

var transactions []transaction

type transaction struct {
	ID              int     `form:"id" json:"id"`
	TransactionCode string  `form:"transaction_code" json:"transaction_code" validate:"required,transaction_code"`
	Currency        string  `form:"currency" json:"currency" validate:"required,currency"`
	Amount          float64 `form:"amount" json:"amount" validate:"required,amount"`
	Receiver        string  `form:"receiver" json:"receiver" validate:"required,receiver"`
	Sender          string  `form:"sender" json:"sender" validate:"required,sender"`
	TransactionDate string  `form:"transaction_date" json:"transaction_date" validate:"required,transaction_date"`
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
	//errs := validateFields(newTransactions)
	errs := url.Values{}

	//tipos := reflect.TypeOf(newTransactions)

	if newTransactions.Receiver == "" {
		errs.Add("Receiver", "The Receiver is required!")
	}
	if newTransactions.Sender == "" {
		errs.Add("Sender", fmt.Sprintf("The Sender is required", reflect.TypeOf(newTransactions)))
	}
	if newTransactions.Amount == 0 {
		errs.Add("Amount", "The Amount is required!")
	}
	if newTransactions.Currency == "" {
		errs.Add("Currency", "The currency is required!")
	}
	if newTransactions.TransactionCode == "" {
		errs.Add("TransactionCode", "The TransactionCode is required!")
	}
	if newTransactions.TransactionDate == "" {
		errs.Add("TransactionDate", "The TransactionDate is required!")
	}
	if len(errs) > 0 {
		context.String(400, "{ \n \"errors\": \"%v\"", errs)
	}

	if err != nil {
		context.String(200, "{ \n \"error\": \"%v\"", err)
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

func validateFields(newTransactions transaction) (errs url.Values) {

	transReflect := reflect.ValueOf(newTransactions)

	for i := 0; i < transReflect.NumField(); i++ {

		varValor := transReflect.Field(i).Interface()
		s := reflect.TypeOf(varValor).Kind()
		field := transReflect.Type().Field(i).Name

		fmt.Println(s)
		fmt.Println(field)
		if fmt.Sprint(s) == "string" {
			if varValor == "" {
				errs.Add(string(field), fmt.Sprintf("The %v is required", transReflect.Type().Field(i).Name))
			}
		} else {
			if varValor == 0.00 {
				errs.Add(string(field), fmt.Sprintf("The %v is required", transReflect.Type().Field(i).Name))
			}
		}
	}

	if newTransactions.Receiver == "" {
		errs.Add("Receiver", "The Receiver is required!")
	}
	if newTransactions.Sender == "" {
		errs.Add("Sender", fmt.Sprintf("The Sender is required", reflect.TypeOf(newTransactions)))
	}
	if newTransactions.Amount == 0 {
		errs.Add("Amount", "The Amount is required!")
	}
	if newTransactions.Currency == "" {
		errs.Add("Currency", "The currency is required!")
	}
	if newTransactions.TransactionCode == "" {
		errs.Add("TransactionCode", "The TransactionCode is required!")
	}
	if newTransactions.TransactionDate == "" {
		errs.Add("TransactionDate", "The TransactionDate is required!")
	}

	return errs
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

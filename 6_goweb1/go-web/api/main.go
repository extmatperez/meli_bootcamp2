package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/sayHi/:name/:lastName", SayHi)

	routerTransactions := router.Group("/transactions")
	{
		routerTransactions.GET("/GetAll", GetAll)
		routerTransactions.GET("/Filter", FilterByParams)
		routerTransactions.GET("/GetByID/:ID", GetByID)
	}

	router.Run(":8080")
}

type Transaction struct {
	ID              int64   `json:"id"`
	TransactionCode string  `json:"transactionCode"`
	Currency        string  `json:"currency"`
	Amount          float64 `json:"amount"`
	Sender          string  `json:"sender"`
	Receiver        string  `json:"receiver"`
	TransactionDate string  `json:"transactionDate"`
}

func GetAll(c *gin.Context) {
	jsonFile, _ := os.Open("transactions.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var transactions []Transaction
	err := json.Unmarshal(byteValue, &transactions)
	if err == nil {
		c.JSON(http.StatusOK, transactions)
	} else {
		c.JSON(http.StatusOK, transactions)
	}
}

func SayHi(c *gin.Context) {
	name := c.Param("name")
	lastName := c.Param("lastName")
	greeting := fmt.Sprintf("Hi, %v %v", name, lastName)
	c.JSON(http.StatusOK, greeting)
}

func ObtainTransactions() []Transaction {
	jsonFile, _ := os.Open("transactions.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var transactions []Transaction
	err := json.Unmarshal(byteValue, &transactions)
	if err != nil {
		return nil
	} else {
		return transactions
	}
}

func GetByID(c *gin.Context) {
	var transactions []Transaction = ObtainTransactions()
	var transactionFinded *Transaction
	if transactions != nil {
		idStr := c.Param("ID")
		id, _ := strconv.ParseInt(idStr, 10, 64)
		for _, transaction := range transactions {
			if transaction.ID == id {
				transactionFinded = &transaction
			}
		}
		if transactionFinded == nil {
			c.JSON(http.StatusNotFound, "Transaction "+idStr+" not found")
		} else {
			c.JSON(http.StatusOK, transactionFinded)
		}
	} else {
		c.JSON(500, nil)
	}
}

func FilterByParams(c *gin.Context) {
	var transactions []Transaction = ObtainTransactions()
	//var filters map[string]interface{} = make(map[string]interface{})
	var transactionsFinded []Transaction

	/*
		Amount          float64 `json:"amount"`
		Sender          string  `json:"sender"`
		Receiver        string  `json:"receiver"`
		TransactionDate
	*/

	for _, tx := range transactions {
		id, err := strconv.ParseInt(c.Query("ID"), 10, 64)
		if err == nil {
			if tx.ID == id {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txCode := c.Query("TransactionCode")
		if txCode != "" {
			if tx.TransactionCode == txCode {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txCurrency := c.Query("Currency")
		if txCurrency != "" {
			if tx.Currency == txCurrency {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
	}

	if len(transactionsFinded) == 0 {
		c.String(404, "No transactions found")
	} else {
		c.JSON(200, &transactionsFinded)
	}
}

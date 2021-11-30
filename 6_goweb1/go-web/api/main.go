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
	ID       int64   `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Remitter string  `json:"remitter"`
	Receptor string  `json:"receptor"`
	Date     string  `json:"date"`
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
	var transactionsFinded []Transaction

	for _, tx := range transactions {
		id, err := strconv.ParseInt(c.Query("ID"), 10, 64)
		if err == nil {
			if !contains(&transactionsFinded, &tx) && tx.ID == id {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txCode := c.Query("Code")
		if txCode != "" {
			if !contains(&transactionsFinded, &tx) && tx.Code == txCode {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txCurrency := c.Query("Currency")
		if txCurrency != "" {
			if !contains(&transactionsFinded, &tx) && tx.Currency == txCurrency {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txAmount, err := strconv.ParseFloat(c.Query("Amount"), 64)
		if err == nil {
			if !contains(&transactionsFinded, &tx) && tx.Amount == txAmount {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txRemitter := c.Query("Remitter")
		if txRemitter != "" {
			if !contains(&transactionsFinded, &tx) && tx.Remitter == txRemitter {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txReceptor := c.Query("Receptor")
		if txReceptor != "" {
			if !contains(&transactionsFinded, &tx) && tx.Receptor == txReceptor {
				transactionsFinded = append(transactionsFinded, tx)
			}
		}
		txDate := c.Query("Date")
		if txDate != "" {
			if !contains(&transactionsFinded, &tx) && tx.Date == txDate {
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

func contains(transactions *[]Transaction, element *Transaction) bool {
	for _, a := range *transactions {
		if a.ID == element.ID {
			return true
		}
	}
	return false
}

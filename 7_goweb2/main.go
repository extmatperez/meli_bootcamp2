package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id              int    `json:"id"`
	TransactionCode string `json:"transactionCode"`
	Currency        string `json:"currency"`
	Amount          int    `json:"amount"`
	Sender          string `json:"sender"`
	Receiver        string `json:"receiver"`
	Date            string `json:"date"`
}

var transactions = []Transaction{}

func getTransactions() ([]byte, error) {
	return os.ReadFile("./transactions.json")
}

func loadData() {
	data, _ := getTransactions()

	json.Unmarshal(data, &transactions)
}

func validatePayload(c *gin.Context, payload Transaction) error {

	r := reflect.ValueOf(payload)

	for i := 0; i < r.NumField(); i++ {
		value := r.Field(i).Interface()
		name := r.Type().Field(i).Name

		fieldType := reflect.TypeOf(value).Kind()
		if strings.ToLower(name) != "id" {

			if fmt.Sprint(fieldType) == "string" {
				if value == "" {
					return fmt.Errorf("%v no puede estar vacio", name)
				}
			}
			if fmt.Sprint(fieldType) == "int" {
				if value == 0 {
					return fmt.Errorf("%v no puede estar vacio", name)
				}
			}
		}
	}
	return nil
}

func createHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		authToken := c.GetHeader("authtoken")

		if authToken != "123" {
			c.JSON(401, gin.H{
				"error": "Token invalido",
			})
			return
		}

		var newTransaction Transaction

		if err := c.ShouldBindJSON(&newTransaction); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := validatePayload(c, newTransaction)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		newTransaction.Id = len(transactions) + 1

		transactions = append(transactions, newTransaction)
		c.JSON(http.StatusCreated, newTransaction)
	}
}

func main() {
	router := gin.Default()

	loadData()

	transactionsRoute := router.Group("/transactions")
	transactionsRoute.POST("/", createHandler())

	router.Run()
}

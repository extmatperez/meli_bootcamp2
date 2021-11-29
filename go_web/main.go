package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID              int    `json:"id"`
	TransactionCode string `json:"transaction_code"`
	Coin            string `json:"coin"`
	Amount          string `json:"amount"`
	Emitor          string `json:"emitor"`
	Receptor        string `json:"receptor"`
	TransactionDate string `json:"transaction_date"`
}

func GetAll(file string) []Transaction {
	data, err := os.ReadFile(file)

	var TransactionList []Transaction

	json.Unmarshal(data, &TransactionList)

	if err != nil {
		fmt.Println("Error json: ", err)
	}
	return TransactionList
}

func main() {

	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola ivan",
		})
	})

	router.GET("/transactions", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetAll("./transactions.json"))
	})

	router.Run(":8080")

}

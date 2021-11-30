package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	//fmt.Printf("%+v\n", transactions)

	router := gin.Default()

	router.GET("/transactions", getAll)
	router.GET("/transactions/:id", getById)

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

	/* content := context.Request.Body
	header := context.Request.Header
	method := context.Request.Method

	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	} */

	transactions := readTransactions()
	context.JSON(200, gin.H{
		"transaction": transactions})
}
func getById(context *gin.Context) {

	transactions := readTransactions()
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

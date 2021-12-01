package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID               int     `form:"id", json:"id"`
	Transaction_Code string  `form:"transaction_code", json:"transaction_code"`
	Coin             string  `form:"coin", json:"coin"`
	Amount           float64 `form:"amount", json:"amount"`
	Emitor           string  `form:"emitor", json:"emitor"`
	Receptor         string  `form:"receptor", json:"receptor"`
	Transaction_Date string  `form:"transaction_date", json:"transaction_date"`
}

func TransactionsGetter() []Transaction {
	data, err := os.ReadFile("./transactions.json")

	var TransactionList []Transaction

	json.Unmarshal(data, &TransactionList)

	if err != nil {
		fmt.Println("Error json: ", err)
	}
	return TransactionList
}

func main() {
	router := gin.Default()

	router.GET("/hola/:nombre", Greetings)
	router.GET("/transactions", GetAll)
	router.GET("/filtrar", FilterQuery)
	router.GET("/transaction/:id", GetOne)

	router.POST("/agregarEntidad", addTransaction)

	router.Run(":8080")

}

func Greetings(ctx *gin.Context) {
	name := ctx.Param("name")

	if name != "" {
		ctx.String(200, "Hola %v!!!", name)
	} else {
		ctx.String(200, "Hola anonimo!")
	}
}
func GetAll(ctx *gin.Context) {
	transactions := TransactionsGetter()
	ctx.JSON(http.StatusOK, transactions)
}

func FilterQuery(ctx *gin.Context) {
	filter := ctx.Query("filter")
	value := ctx.Query("value")
	//header := ctx.Request.Body

	//for k, v := range header {
	fmt.Printf("%v\n", filter)
	fmt.Printf("%v\n", value)

	//}

	var TransactionList []Transaction

	switch filter {
	case "transaction_code":
		TransactionList = SearchByTransactionCode(value)
		ctx.JSON(http.StatusOK, TransactionList)
	case "coin":
		TransactionList = SearchByCoin(value)
		ctx.JSON(http.StatusOK, TransactionList)
	case "amount":
		TransactionList = SearchInRangeAmount(value)
		ctx.JSON(http.StatusOK, TransactionList)
	case "emitor":
		TransactionList = SearchEmitor(value)
		ctx.JSON(http.StatusOK, TransactionList)
	case "receptor":
		TransactionList = SearchReceptor(value)
		ctx.JSON(http.StatusOK, TransactionList)
	case "transaction_date":
		TransactionList = SearchDate(value)
		ctx.JSON(http.StatusOK, TransactionList)
	default:
		ctx.String(400, "Invalid Query")
	}
}

func SearchReceptor(receptor string) []Transaction {
	var TransactionList []Transaction = TransactionsGetter()

	var transactionToReceptor []Transaction

	for _, tx := range TransactionList {
		if receptor == tx.Receptor {
			transactionToReceptor = append(transactionToReceptor, tx)
		}
	}

	return transactionToReceptor
}

func SearchByCoin(coin string) []Transaction {
	var TransactionList []Transaction = TransactionsGetter()

	var transactionsByCoin []Transaction

	for _, tx := range TransactionList {
		if coin == tx.Coin {
			transactionsByCoin = append(transactionsByCoin, tx)
		}
	}

	return transactionsByCoin
}

func SearchEmitor(emitor string) []Transaction {
	var TransactionList []Transaction = TransactionsGetter()

	var transactionFromEmitor []Transaction

	for _, tx := range TransactionList {
		if emitor == tx.Emitor {
			transactionFromEmitor = append(transactionFromEmitor, tx)
		}
	}

	return transactionFromEmitor
}

func SearchInRangeAmount(amount string) []Transaction {
	amountFloat, _ := strconv.ParseFloat(amount, 64)
	var TransactionList []Transaction = TransactionsGetter()

	var transactionsBetweenAmount []Transaction

	for _, tx := range TransactionList {
		if amountFloat == tx.Amount {
			transactionsBetweenAmount = append(transactionsBetweenAmount, tx)
		}
	}

	return transactionsBetweenAmount
}

func SearchDate(date string) []Transaction {
	var TransactionList []Transaction = TransactionsGetter()

	var transactionsByDate []Transaction

	for _, tx := range TransactionList {
		if date == tx.Transaction_Date {
			transactionsByDate = append(transactionsByDate, tx)
		}
	}

	return transactionsByDate
}

func SearchByTransactionCode(code string) []Transaction {
	var TransactionList []Transaction = TransactionsGetter()

	var transactionsByCode []Transaction

	for _, tx := range TransactionList {
		if code == tx.Transaction_Code {
			transactionsByCode = append(transactionsByCode, tx)
		}
	}

	return transactionsByCode
}

func GetOne(ctx *gin.Context) {
	var TransactionList []Transaction = TransactionsGetter()

	var TransactionFound Transaction

	idSearched, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.String(500, "Error")
	}

	finded := false

	for _, tx := range TransactionList {
		if tx.ID == idSearched {
			TransactionFound = tx
			finded = true
			break
		}
	}

	if finded {
		ctx.JSON(http.StatusOK, TransactionFound)
	} else {
		ctx.String(400, "No se encontro una transaccion con el id: %v", idSearched)
	}
}

////////////////////////////////////////////////////////////////

var ListaTransactions []Transaction

func addTransaction(ctx *gin.Context) {
	var transaction Transaction

	token := ctx.GetHeader("token")

	if token == "123456" {
		err := ctx.ShouldBindJSON(&transaction)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			validFields := fieldsValidator(transaction)
			if validFields == "ok" {
				var id = len(ListaTransactions) + 1
				transaction.ID = id
				ListaTransactions = append(ListaTransactions, transaction)
				ctx.JSON(200, transaction)
			} else {
				ctx.String(400, "el campo %s es requerido", validFields)
			}
		}
	} else {
		ctx.String(401, "no tiene permisos para realizar la petición solicitada")
	}
}

func fieldsValidator(transaction Transaction) string {
	typesOfFields := reflect.ValueOf(transaction)

	i := 0
	for i = 0; i < typesOfFields.NumField(); i++ {
		valueOfField := typesOfFields.Field(i).Interface()
		if valueOfField == "" {
			field := reflect.TypeOf(transaction).Field(i).Name
			return field
		}
	}
	return "ok"
}

package main

import (
	"github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/cmd/server/handler"
	transactions "github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/internal/transactions"
	stores "github.com/extmatperez/meli_bootcamp2/tree/ghione_andres/7_goweb2/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//api.Use()
	//router.GET("/sayHi/:name/:lastName", SayHi)
	db := stores.New(stores.FileType, "/Users/aghione/Desktop/repositorios/bootcamp/practicas/meli_bootcamp2/7_goweb2/internal/transactions/transactions.json")
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	controller := handler.NewTransaction(service)

	routerTransactions := router.Group("/transactions")
	{
		routerTransactions.GET("/GetAll", controller.GetAll())
		//routerTransactions.GET("/Filter", FilterByParams)
		routerTransactions.GET("/GetByID/:ID", controller.GetByID())
		routerTransactions.POST("/", controller.Store())
		//routerTransactions.GET("/chargeData", ObtainTransactions)
		routerTransactions.PUT("/:ID", controller.Update())
		routerTransactions.DELETE("/:ID", controller.Delete())
		routerTransactions.PATCH("/transactionCode/:ID/:TransactionCode", controller.ModifyTransactionCode())
		routerTransactions.PATCH("/amount/:ID/:Amount", controller.ModifyAmount())
	}

	router.Run(":8080")
}

/*
func AddTransaction(context *gin.Context) {
	fmt.Println(transactions)
	var newTransaction Transaction
	err := context.ShouldBindJSON(&newTransaction)
	if err == nil {
		if ValidateTransaction(newTransaction) {
			//agregar al archivo
			if len(transactions) == 0 {
				newTransaction.ID = 1
			} else {
				newTransaction.ID = (transactions)[len(transactions)-1].ID+1
			}
			transactions = append(transactions, newTransaction)
			context.JSON(http.StatusOK, newTransaction)
			return
		} else {
			context.String(http.StatusBadRequest, fmt.Sprintf("Invalid fields in transaction %v", newTransaction))
		}
	} else {
		context.String(http.StatusInternalServerError, fmt.Sprintf("Exploded"))
	}
}

func ValidateTransaction(transaction Transaction) bool {
	validTransaction := false
	types := reflect.TypeOf(transaction)
	for i := 0; i < types.NumField(); i++ {
		if types.Field(i).Name != "ID" && types.Field(i).Type.String() == "string" {
			if reflect.ValueOf(transaction).Field(i).Interface() != "" {
				validTransaction = true
			} else {
				return false
			}
		} else if types.Field(i).Type.String() == "int64" {
			if reflect.ValueOf(transaction).Field(i).Int() > 0 {
				validTransaction = true
			} else {
				return false
			}
		} else if types.Field(i).Type.String() == "float64" {
			if reflect.ValueOf(transaction).Field(i).Float() > 0 {
				validTransaction = true
			} else {
				return false
			}

		}
	}
	return validTransaction
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

func ObtainTransactions(c *gin.Context) {
	jsonFile, _ := os.Open("transactions.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err := json.Unmarshal(byteValue, &transactions)
	if err != nil {
		transactions = nil
	}
}

func GetByID(c *gin.Context) {
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
	var transactionsFinded []Transaction

	for _, tx := range transactions {
		id, errId := strconv.ParseInt(c.Query("ID"), 10, 64)
		if errId == nil {
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
		txAmount, errAmount := strconv.ParseFloat(c.Query("Amount"), 64)
		if errAmount == nil {
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
		if errId != nil && txCode == "" && txCurrency == "" && errAmount != nil && txRemitter == "" && txReceptor == "" && txDate == "" {
			transactionsFinded = transactions
		}
	}

	if len(transactionsFinded) == 0 {
		c.String(http.StatusNotFound, "No transactions found")
	} else {
		c.JSON(http.StatusOK, &transactionsFinded)
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
*/

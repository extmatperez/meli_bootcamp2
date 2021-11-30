package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var transacciones []Transaction

type Transaction struct {
	ID                  int     `json:"id"`
	CodigoDeTransaccion string  `json:"codigo_de_transaccion" binding:"required"`
	Moneda              string  `json:"moneda" binding:"required"`
	Monto               float64 `json:"monto" binding:"required"`
	Emisor              string  `json:"emisor" binding:"required"`
	Receptor            string  `json:"receptor" binding:"required"`
	FechaDeTransaccion  string  `json:"fecha_de_transaccion" binding:"required"`
}

//moneda, emisor, receptor, fecha,>>.. eur , mxn
func sayHello(c *gin.Context) {
	var status int
	message := ""
	messageName := "message"

	if name, ok := c.GetQuery("nombre"); ok {
		if name != "" {
			status = 200
			message = "¡Hola, " + name + "!"

		} else {
			status = 400
			message = "'nombre' cannot be empty"
			messageName = "error"
		}
	} else {
		status = 400
		message = "param 'nombre' not received"
		messageName = "error"
	}

	c.JSON(status, gin.H{
		messageName: message,
	})

}

func ReadTransactionsFile() (status int, response []Transaction, err error) {
	data, err := os.ReadFile("./transactions.json")

	if err != nil {
		status = http.StatusInternalServerError
	} else {
		err = json.Unmarshal(data, &response)
		if err != nil {
			status = http.StatusInternalServerError
		} else {
			status = http.StatusOK
		}
	}
	return
}

func FiltrarTransacciones(moneda string, data []Transaction, ctx *gin.Context) (transacciones []Transaction, err error) {
	for _, t := range data {
		if t.Moneda == moneda {
			transacciones = append(transacciones, t)
		}
	}
	return transacciones, nil
}

func GetAll(c *gin.Context) {

	//data, err := os.ReadFile("./transactions.json")
	response := transacciones
	var status int
	//c.JSON(status, response)
	filter := c.Query("moneda")
	status = http.StatusOK
	if filter != "" {
		response, _ = FiltrarTransacciones(filter, response, c)
		//status = http.StatusOK
	}
	if len(response) == 0 {
		//status = http.StatusOK
	}
	c.JSON(status, response)
}

func GetTransactionByID(ctx *gin.Context) {
	found := false

	for _, tx := range transacciones {
		if ctx.Param("id") == strconv.Itoa(tx.ID) {
			ctx.JSON(http.StatusOK, tx)
			found = true
			break
		}
	}
	if !found {
		ctx.JSON(http.StatusNotFound, "No transactions match the id "+ctx.Param("id"))
	}

}
func CreateTransaction(ctx *gin.Context) {

	if ValidateToken(ctx) {
		var tr Transaction

		err := ctx.ShouldBindJSON(&tr)

		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if len(transacciones) == 0 {
			tr.ID = 1
		} else {
			tr.ID = transacciones[len(transacciones)-1].ID + 1
		}
		transacciones = append(transacciones, tr)
		ctx.JSON(http.StatusCreated, tr)
	}

}

func ValidateToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	if token != "" {
		if token == "123456" {
			return true
		}
		ctx.String(http.StatusUnauthorized, "Token incorrecto")
		return false
	}
	ctx.String(http.StatusUnauthorized, "No se ingresó un token")
	return false
}

func main() {

	data, _ := os.ReadFile("./transactions.json")
	json.Unmarshal(data, &transacciones)

	router := gin.Default()

	router.GET("/hello", sayHello)
	//router.GET("/transactions", GetAll)
	transac := router.Group("/transactions")
	{
		transac.GET("", GetAll)
		transac.GET("/:id", GetTransactionByID)
		transac.POST("/add", CreateTransaction)
	}

	router.Run()

}

package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	ID                  int     `json:"id"`
	CodigoDeTransaccion string  `json:"codigo_de_transaccion"`
	Moneda              string  `json:"moneda"`
	Monto               float64 `json:"monto"`
	Emisor              string  `json:"emisor"`
	Receptor            string  `json:"receptor"`
	FechaDeTransaccion  string  `json:"fecha_de_transaccion"`
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

	status, response, err := ReadTransactionsFile()

	//data, err := os.ReadFile("./transactions.json")

	if err != nil {
		c.JSON(status, gin.H{
			"error": "could not found transactions",
		})

	} else { //pude leer el archivo correctamente
		//c.JSON(status, response)
		filter := c.Query("moneda")

		if filter != "" {
			response, _ = FiltrarTransacciones(filter, response, c)
		}
		if len(response) == 0 {
			status = http.StatusNotFound
		}
		c.JSON(status, response)
	}

}

func GetTransactionByID(ctx *gin.Context) {
	status, response, err := ReadTransactionsFile()
	found := false
	if err != nil {
		ctx.JSON(status, "Unable to get transactions")
	} else {

		for _, tx := range response {
			if ctx.Param("id") == strconv.Itoa(tx.ID) {
				ctx.JSON(status, tx)
				found = true
				break
			}
		}
		if !found {
			ctx.JSON(http.StatusNotFound, "No transactions match the id "+ctx.Param("id"))
		}
	}
}

func main() {

	router := gin.Default()

	router.GET("/hello", sayHello)
	//router.GET("/transactions", GetAll)
	transac := router.Group("/transactions")
	{
		transac.GET("", GetAll)
		transac.GET("/:id", GetTransactionByID)
	}

	router.Run()

}

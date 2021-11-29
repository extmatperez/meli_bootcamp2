package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


 
type Transaccion []struct {
	ID       int    `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}


func GetTransactionFromFolder(fileName string) Transaccion{
	file, _ := ioutil.ReadFile(fileName)
	
	var transaction Transaccion
 
	_ = json.Unmarshal([]byte(file), &transaction)

	return transaction

}

func GetAllTransactions(router *gin.Engine,filename string){
	transactions := GetTransactionFromFolder(filename)

	router.GET("/alltransaction", func(c *gin.Context){
		c.JSON(http.StatusOK,transactions)
	})
}
func main() {
	ruta := "./6_goweb1/transactions.json"

	router := gin.Default()
	GetAllTransactions(router,ruta)
	router.Run()

 }



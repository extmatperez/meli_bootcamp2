package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	ID       int    `json:"id"`
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}


func GetTransactionFromFolder(fileName string) ([]Transaccion,error){
	file, _ := ioutil.ReadFile(fileName)
	
	var transaction []Transaccion
 
	err := json.Unmarshal([]byte(file), &transaction)
	
	if(err != nil) {
		return nil,err
	}
	return transaction,nil

}

func GetAllTransactions(c *gin.Context){
	filename := "./6_goweb1/transactions.json"
	transactions,err := GetTransactionFromFolder(filename)

		fmt.Print(c)
		if(err != nil){
		 c.String(http.StatusForbidden,"No hay datos en el filename: "+filename)
		}else{
     	 c.JSON(http.StatusOK,transactions)
		}
	

}
func main() {

	router := gin.Default()
	router.GET("/alltransaction",GetAllTransactions)
	router.Run()

 }



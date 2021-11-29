package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	ID       int   `json:"id"`
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


func GetTransactionById(c *gin.Context){
	filename := "./6_goweb1/transactions.json"
	transactions,err := GetTransactionFromFolder(filename)

		fmt.Print(c)
		if(err != nil){
		 c.String(http.StatusForbidden,"No hay datos en el filename: "+filename)
		}else{
			idQuery,_ := strconv.Atoi(c.Param("id"))
			var transacccion *Transaccion
			encontre := false
			for _,v := range transactions{
				if  idQuery== v.ID{			
					encontre = true	
					transacccion = &v
					break
				}
				
			}

			if(encontre){
				c.JSON(http.StatusOK,transacccion)
			}else{
				c.String(http.StatusForbidden,"No existe la transaccion con el id ingresado")
			}
		}
		



}


func FindInclusive(c *gin.Context){
		filename := "./6_goweb1/transactions.json"
		transactions,err := GetTransactionFromFolder(filename)

		fmt.Print(c)
		if(err != nil){
		 c.String(http.StatusForbidden,"No hay datos en el filename: "+filename)
		}
     	var transaction Transaccion
		body := c.BindJSON(&transaction)
		
		if(body != nil){
			c.String(http.StatusForbidden,"Debes pasar un json con los datos a buscar")
		   }

		 var filtrados []Transaccion  
		 for _,v := range transactions{
			if(v.Codigo == transaction.Codigo || v.Emisor == transaction.Emisor || v.Fecha == transaction.Fecha||
				v.Moneda == transaction.Moneda|| v.Monto == transaction.Monto || v.Receptor == transaction.Receptor){
				filtrados = append(filtrados, v)
			}
		    }

			if len(filtrados) == 0{
				c.String(http.StatusForbidden,"No hay datos con esos filtros")
			}else{
				c.JSON(http.StatusOK,filtrados)
			}
}




func main() {

	router := gin.Default()
	router.GET("/alltransaction",GetAllTransactions)
	router.GET("/findinclusive",FindInclusive)
	router.GET("/transaction/:id",GetTransactionById)
	router.Run()

 }



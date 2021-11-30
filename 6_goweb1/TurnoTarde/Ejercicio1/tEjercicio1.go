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


func FindExlusive(c *gin.Context){
	filename := "./6_goweb1/transactions.json"
	transactions,err := GetTransactionFromFolder(filename)

	if(err != nil){
	 c.String(http.StatusForbidden,"No hay datos en el filename: "+filename)
	}
	 var parametros Transaccion
	 body := c.BindJSON(&parametros)
	 fmt.Println(parametros)
	if(body != nil){
		c.String(http.StatusForbidden,"Debes pasar un json con los datos a buscar")
	   }

	   var filtrados []Transaccion  
	   filtros := GetFiltros(parametros)
	   

	if(len(filtros) == 0){
		c.String(http.StatusForbidden,"Debes pasar al menos un flitro con los datos a buscar")
	}

		for _,v := range transactions{
			flag:=false
			numFiltros:=0
			for _,f := range filtros{
			
				if(f == "Codigo" && parametros.Codigo != "" && v.Codigo == parametros.Codigo){
					flag = true
					numFiltros++
					fmt.Println("Codigo",v.Codigo)
				}else if(f == "Emisor" && parametros.Emisor != "" && v.Emisor == parametros.Emisor ){
					flag = true
					numFiltros++
					fmt.Println("Emisor",v.Codigo)
				}else if(f == "Fecha" && parametros.Fecha != "" && v.Fecha == parametros.Fecha){
					flag = true
				}else if(f == "Moneda" && parametros.Moneda != "" && v.Moneda == parametros.Moneda){
					flag = true
					numFiltros++
					fmt.Println("Moneda",v.Codigo)
				}else if(f == "Monto" && parametros.Monto!= "" && v.Monto == parametros.Monto){
					flag = true
					numFiltros++
				}else if(f == "Receptor" && parametros.Receptor != "" && v.Receptor == parametros.Receptor){
					flag = true
					numFiltros++
				}else {
					flag = false
					fmt.Println("Ninguno",v.Codigo)
				}

			}			
			if flag && numFiltros == len(filtros) {
				fmt.Println("Print",v.Codigo)
				filtrados = append(filtrados, v)
			}
		
		}

	c.JSON(http.StatusOK,filtrados)
		
}

func GetFiltros(parametros Transaccion) []string{
	var list []string
	if(parametros.Codigo != ""){
		list = append(list, "Codigo")
	}
	if(parametros.Emisor != ""){
		list = append(list, "Emisor")
	}

	if(parametros.Fecha != ""){
		list = append(list, "Fecha")
	}

	if(parametros.Moneda != ""){
		list = append(list, "Moneda")
	}

	if(parametros.Monto != ""){
		list = append(list, "Monto")
	}

	if(parametros.Receptor != ""){
		list = append(list, "Receptor")
	}


	return list;


}



func main() {

	server := gin.Default()

	transaction := server.Group("/transactions")
	{
		transaction.GET("/findinclusive",FindInclusive)
		transaction.GET("/findexclusive",FindExlusive)
		transaction.GET("/:id",GetTransactionById)
		transaction.GET("/all",GetAllTransactions)
	}

	
	server.Run()

 }



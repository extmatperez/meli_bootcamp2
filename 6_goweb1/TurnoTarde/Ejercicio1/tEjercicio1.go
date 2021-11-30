package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const(
	Codigo = "Codigo"
	Moneda = "Moneda"
	Monto = "Monto"
	Emisor = "Emisor"
	Receptor = "Receptor"
	Fecha = "Fecha"
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


func GetTransactionFromFolder() ([]Transaccion,error){
	fileName := "./6_goweb1/transactions.json"
	file, err1:= ioutil.ReadFile(fileName)
	if(err1 != nil) {
		return nil,err1
	}

	var transaction []Transaccion
 
	err := json.Unmarshal([]byte(file), &transaction)
	
	if(err != nil) {
		return nil,err
	}


	return transaction,nil

}

func GetAllTransactions(c *gin.Context){
	transactions,err := GetTransactionFromFolder()

		fmt.Print(c)
		if(err != nil){
		 c.String(http.StatusForbidden,"No hay datos en el filename.",err.Error())
		}
     	 c.JSON(http.StatusOK,transactions)
		
}


func GetTransactionById(c *gin.Context){
	transactions,err := GetTransactionFromFolder()

		fmt.Print(c)
		if(err != nil){
		c.String(http.StatusForbidden,"No hay datos en el filename.",err.Error())
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


func GetTransactionsInclusive(c *gin.Context){
		transactions,err := GetTransactionFromFolder()
		if(err != nil){
			c.String(http.StatusForbidden,"No hay datos en el filename.",err.Error())
		}

     	var transaction Transaccion
		err2 := c.BindJSON(&transaction)
		if(err2 != nil){
			c.String(http.StatusForbidden,"Debes pasar un json con los datos a buscar. Error: ",err2.Error())
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


func GetTransactionsExlusive(c *gin.Context){
		
		transactions,err := GetTransactionFromFolder()
		if(err != nil){
			c.String(http.StatusForbidden,"No hay datos en el filename.",err.Error())
		}

		var parametros Transaccion
		err1 := c.BindJSON(&parametros)
		fmt.Println(parametros)
			if(err1 != nil){
			c.String(http.StatusForbidden,"Debes pasar un json con los datos a buscar")
		}

	   
	   filtros := GetFiltros(parametros)
	   if(len(filtros) == 0){
			c.String(http.StatusForbidden,"Debes pasar al menos un flitro con los datos a buscar")
		}
		 filtrados := transactions

		for _,filtro := range filtros{
			fmt.Println("filtro",filtro)
			fmt.Println("valor", reflect.ValueOf(parametros).FieldByName(filtro).String())
		
			filtrados = filtrar(filtrados,filtro,reflect.ValueOf(parametros).FieldByName(filtro).String())
		}
		

		c.JSON(http.StatusOK,filtrados)
		
}


func GetFiltrados(filtros []string,transactions []Transaccion ,parametros Transaccion) []Transaccion {
	var filtrados []Transaccion 

		for _,v := range transactions{
			flag:=false
			numFiltros:=0
			for _,f := range filtros{
			
				if(f == Codigo && parametros.Codigo != "" && v.Codigo == parametros.Codigo){
					flag = true
					numFiltros++
				}else if(f == Emisor && parametros.Emisor != "" && v.Emisor == parametros.Emisor ){
					flag = true
					numFiltros++
				}else if(f == Fecha && parametros.Fecha != "" && v.Fecha == parametros.Fecha){
					flag = true
				}else if(f == Moneda && parametros.Moneda != "" && v.Moneda == parametros.Moneda){
					flag = true
					numFiltros++
				}else if(f == Monto && parametros.Monto!= "" && v.Monto == parametros.Monto){
					flag = true
					numFiltros++
				}else if(f == Receptor && parametros.Receptor != "" && v.Receptor == parametros.Receptor){
					flag = true
					numFiltros++
				}else {
					flag = false
				}

			}			
			if flag && numFiltros == len(filtros) {
				fmt.Println("Print",v.Codigo)
				filtrados = append(filtrados, v)
			}
		}
		return filtrados
}



func filtrar(sliceTransaccion[]Transaccion, campo string, valor string) []Transaccion {
	var filtrado []Transaccion

	var per Transaccion
	tipos := reflect.TypeOf(per)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		// fmt.Println(i, "->", tipos.Field(i).Name)
		if tipos.Field(i).Name == campo {
			break
		}
	}

	for _, v := range sliceTransaccion {
		cadena := fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor) {
			// if reflect.ValueOf(v).Field(i).Interface() == valor {
			filtrado = append(filtrado, v)
		}
	}

	return filtrado
}








func GetFiltros(parametros Transaccion) []string{
	var list []string
	if(parametros.Codigo != ""){
		list = append(list, Codigo)
	}
	if(parametros.Emisor != ""){
		list = append(list, Emisor)
	}

	if(parametros.Fecha != ""){
		list = append(list, Fecha)
	}

	if(parametros.Moneda != ""){
		list = append(list, Moneda)
	}

	if(parametros.Monto != ""){
		list = append(list, Monto)
	}

	if(parametros.Receptor != ""){
		list = append(list, Receptor)
	}

	return list;
}



func main() {

	server := gin.Default()

	transaction := server.Group("/transactions")
	{

		transaction.GET("/",GetAllTransactions)
		transaction.GET("/:id",GetTransactionById)
		transaction.GET("/findinclusive",GetTransactionsInclusive)
		transaction.GET("/findexclusive",GetTransactionsExlusive)
		
	}

	
	server.Run()

 }



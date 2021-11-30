package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"github.com/gin-gonic/gin"
)

var fileName = "./transactions.json"

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



func filtrar(sliceTransaccion[]Transaccion, campo string, valor string) []Transaccion {
	var filtrado []Transaccion

	var per Transaccion
	tipos := reflect.TypeOf(per)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		if tipos.Field(i).Name == campo {
			break
		}
	}

	for _, v := range sliceTransaccion {
		cadena := fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		 if cadena == valor {
			filtrado = append(filtrado, v)
		}
	}

	return filtrado
}


func GetFiltros(parametros Transaccion) []string{
	var list []string
	r := reflect.ValueOf(parametros)

	for i := 0; i < r.NumField(); i++ {
		varValor := r.Field(i).Interface()
		fmt.Println(varValor)
		if(varValor != "" && varValor != 0){

			list = append(list, r.Type().Field(i).Name)
		}

	}
	return list;
}

func ValidarParametros(parametros Transaccion) []string{
	var list []string
	r := reflect.ValueOf(parametros)

	for i := 0; i < r.NumField(); i++ {
		varValor := r.Field(i).Interface()
		fmt.Println(varValor)
		if(varValor == "" || varValor == 0){
			if(r.Type().Field(i).Name != "ID"){
				list = append(list, r.Type().Field(i).Name)
			}
			
			
		}

	}
	return list;
	
}

func ValidarToken(token string) error{
	if token == "" {
		return errors.New("token Vacio")
	}
	if token != "1234" {
		return errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return nil
}


func InsertTransaction(c *gin.Context){

	var tran Transaccion
	err1 := c.ShouldBindJSON(&tran)

	if err1 != nil {
	c.String(http.StatusBadRequest, "Se produjo un error: %v", err1.Error())
		return
	}

	errs := ValidarToken(c.GetHeader("token"))
	if  errs != nil {
		c.String(http.StatusUnauthorized,errs.Error())
		return
	}

	validar := ValidarParametros(tran)
	if(len(validar) > 0){
		 c.String(http.StatusBadRequest, "Faltan los campos %v", validar)
		 return
		}

	transactions,err2 := GetTransactionFromFolder()

	if(err2 != nil){
		c.String(http.StatusForbidden,"No hay datos en el filename.",err2.Error())
		return
	}

	if len(transactions) == 0 {
		tran.ID = 1
	} else {
		tran.ID = transactions[len(transactions)-1].ID + 1
	}
	
	transactions = append(transactions, tran)

	dataBytes, err3 := json.Marshal(transactions)
    if err3 != nil {
		c.String(http.StatusForbidden,"Error convertir a json.",err3.Error())
		return
    }


	err4 := ioutil.WriteFile(fileName, dataBytes, 0644)
	if err4 != nil {
		c.String(http.StatusForbidden,"Error al guardar datos.",err4.Error())
		return
    }
	
	c.JSON(http.StatusOK, tran)

}




func main() {

	server := gin.Default()

	transaction := server.Group("/transactions")
	{

		transaction.GET("/",GetAllTransactions)
		transaction.GET("/:id",GetTransactionById)
		transaction.GET("/findinclusive",GetTransactionsInclusive)
		transaction.GET("/findexclusive",GetTransactionsExlusive)
		transaction.POST("/",InsertTransaction)
		
	}

	
	server.Run()

 }



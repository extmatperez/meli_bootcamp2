package handler

import (
	
	"net/http"
	"reflect"
	"strconv"
	tran "github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/7_goweb2/TurnoTarde/internal/transaccion"
	"github.com/gin-gonic/gin"
)

type request struct {
	Codigo   string `json:"codigo"`
	Moneda   string `json:"moneda"`
	Monto    string `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

type Transaction struct {
	service tran.Service
}

func NewTransaction(service tran.Service) *Transaction{
	return &Transaction{service}
}

func (tran Transaction) GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		transactions, err := tran.service.GetAll()

		if err != nil {
			ctx.String(http.StatusBadRequest, "Hubo un error %v", err)
		} else {
			ctx.JSON(http.StatusOK, transactions)
		}
	}
}

func (tran Transaction) Store() gin.HandlerFunc{
	return func(c *gin.Context) {
		var transactionRecived request
		
		err := c.ShouldBindJSON(&transactionRecived)

		if err != nil {
			c.String(http.StatusBadRequest, "Hubo un error %v", err)
			return
		} 

		invalidParams := InValidParams(transactionRecived)

		// validamos que esten todos los parametros en el request
		if(len(invalidParams) > 0){
			c.String(http.StatusBadRequest, "Faltan los campos %v", invalidParams)
			return
		}


		tranUpdate,err := tran.service.Store(transactionRecived.Codigo,transactionRecived.Moneda,transactionRecived.Monto,
							transactionRecived.Emisor,transactionRecived.Receptor,transactionRecived.Fecha)

		
		if err != nil {
			c.String(http.StatusBadRequest, "Hubo un error %v", err)
			return
		} 

		c.JSON(http.StatusOK,tranUpdate)

	}
}


func (tran Transaction) Update() gin.HandlerFunc{
	return func(c *gin.Context) {
		idTransaction,err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.String(http.StatusBadRequest, "Hubo un error %v", err)
			return
		} 

		var transactionRecived request
		
		err = c.ShouldBindJSON(&transactionRecived)

		if err != nil {
			c.String(http.StatusBadRequest, "Hubo un error %v", err)
			return
		} 

		invalidParams := InValidParams(transactionRecived)

		// validamos que esten todos los parametros en el request
		if(len(invalidParams) > 0){
			c.String(http.StatusBadRequest, "Faltan los campos %v", invalidParams)
			return
		}


		tranUpdate,err := tran.service.Update(idTransaction,transactionRecived.Codigo,transactionRecived.Moneda,transactionRecived.Monto,
							transactionRecived.Emisor,transactionRecived.Receptor,transactionRecived.Fecha)

		
		if err != nil {
			c.String(http.StatusBadRequest, "Hubo un error %v", err)
			return
		} 

		c.JSON(http.StatusOK,tranUpdate)

	}
}


func (tran Transaction) UpdateCodigoAndMonto() gin.HandlerFunc{
	return func(c *gin.Context) {

			idTransaction,err := strconv.Atoi(c.Param("id"))
	
			if err != nil {
				c.String(http.StatusBadRequest, "Hubo un error %v", err)
				return
			} 
	
			var transactionRecived request
			
			err = c.ShouldBindJSON(&transactionRecived)
	
			if err != nil {
				c.String(http.StatusBadRequest, "Hubo un error %v", err)
				return
			} 

			// validamos que esten todos los monto y codigo en el request
			if(transactionRecived.Codigo == "" || transactionRecived.Monto == "") {
				c.String(http.StatusBadRequest, "mal parametros")
				return
			}

			tranUpdate,err := tran.service.UpdateCodigoAndMonto(idTransaction,transactionRecived.Codigo,
																transactionRecived.Monto)
	
			
			if err != nil {
				c.String(http.StatusBadRequest, "Hubo un error %v", err)
				return
			} 
	
			c.JSON(http.StatusOK,tranUpdate)
	
		}
}

func (tran Transaction) Delete() gin.HandlerFunc{
	return func(c *gin.Context) {
		idTransaction,err := strconv.Atoi(c.Param("id"))
	
			if err != nil {
				c.String(http.StatusBadRequest, "Hubo un error %v", err)
				return
			} 

			err = tran.service.Delete(idTransaction)

			if err != nil {
				c.String(http.StatusBadRequest, "Hubo un error %v", err)
				return
			} 

			c.String(http.StatusOK,"Se elimino correctamente la transaccion con id: %v",idTransaction)

	}
}









func InValidParams(parametros request) []string{
	var list []string
	r := reflect.ValueOf(parametros)

	for i := 0; i < r.NumField(); i++ {
		varValor := r.Field(i).Interface()
		if(varValor == "" || varValor == 0){
			list = append(list, r.Type().Field(i).Name)
						
		}

	}
	return list;
	
}
package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	tra "github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/internal/transaccion"
	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/8_goweb3/pkg/store/web"
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
	service tra.Service
}

func NewTransaction(service tra.Service) *Transaction{
	return &Transaction{service}
}


// ListTransactions godoc
// @Summary List Transactions 
// @Tags Transactions 
// @Description get Transactions 
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Error 401 {object} web.Response
// @Router /transactions [get]
func (tran Transaction) GetAll() gin.HandlerFunc{
	return func(c *gin.Context) {

		err := ValidateToken(c.GetHeader("token"))
		if  err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized,nil,err.Error()))
			return
		}


		transactions, err := tran.service.GetAll()

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
		} else {
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK,transactions,""))
		}
	}
}

// TransactionById godoc
// @Summary Transaction by Id 
// @Tags Transactions 
// @Description get transactions by id
// @Produce json
// @Param token header string true "token"
// @Param id query integer true "id"
// @Success 200 {object} web.Response
// @Error 401 {object} web.Response
// @Router /transactions:id [get]
func (tran Transaction) GetTransactionById() gin.HandlerFunc{
	return func(c *gin.Context) {

		err := ValidateToken(c.GetHeader("token"))
		if  err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized,nil,err.Error()))
			return
		}


		idTransaction,err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 
		transaction, err := tran.service.GetTransactionById(idTransaction)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 
			
		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK,transaction,""))
		
	}
}


// FilterTransactions godoc
// @Summary Filter Transactions  
// @Tags Transactions 
// @Description get transactions by filter from json
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "filtros"
// @Success 200 {object} web.Response
// @Error 401 {object} web.Response
// @Router /transactions/filtros [get]
func (tran Transaction) GetTransactionsExlusive() gin.HandlerFunc{
	return func(c *gin.Context) {

		err := ValidateToken(c.GetHeader("token"))
		if  err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized,nil,err.Error()))
			return
		}

		transactions, err := tran.service.GetAll()

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 

		var parametros request
		err = c.BindJSON(&parametros)
			if(err != nil){
				c.JSON(http.StatusForbidden, web.NewResponse(http.StatusForbidden,nil,err.Error()))
			return
		}

	   
	   filtros := GetParamsFromBody(parametros)

	   if(len(filtros) == 0){
		c.JSON(http.StatusForbidden, web.NewResponse(http.StatusForbidden,nil,
													"Debes pasar al menos un flitro con los datos a buscar"))
			return
		}

		filtrados := transactions

		for _,filtro := range filtros{
			fmt.Println(filtro)
			fmt.Println(reflect.ValueOf(parametros).FieldByName(filtro).String())
			filtrados = filtrar(filtrados,filtro,reflect.ValueOf(parametros).FieldByName(filtro).String())
		}		

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK,filtrados,""))


	}
}

// StoreTransactions godoc
// @Summary Store Transactions 
// @Tags Transactions 
// @Description store Transactions 
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param transaction body request true "Transactions  to store"
// @Success 200 {object} web.Response
// @Router /transactions  [post]
func (tran Transaction) Store() gin.HandlerFunc{
	return func(c *gin.Context) {

		err := ValidateToken(c.GetHeader("token"))
		if  err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized,nil,err.Error()))
			return
		}
		var transactionRecived request
		
		err = c.ShouldBindJSON(&transactionRecived)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 

		invalidParams := InValidParams(transactionRecived)

		// validamos que esten todos los parametros en el request
		if(len(invalidParams) > 0){
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,fmt.Sprintf("Faltan los campos %v", invalidParams)))
			return
		}


		tranUpdate,err := tran.service.Store(transactionRecived.Codigo,transactionRecived.Moneda,transactionRecived.Monto,
							transactionRecived.Emisor,transactionRecived.Receptor,transactionRecived.Fecha)

		
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK,tranUpdate,""))

	}
}

// UpdateTransactions godoc
// @Summary Update Transactions 
// @Tags Transactions 
// @Description store Transactions 
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id query integer true "id"
// @Param transaction body request true "Transactions  to store"
// @Success 200 {object} web.Response
// @Router /transactions:id  [put]
func (tran Transaction) Update() gin.HandlerFunc{
	return func(c *gin.Context) {

		err := ValidateToken(c.GetHeader("token"))
		if  err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized,nil,err.Error()))
			return
		}



		idTransaction,err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 

		var transactionRecived request
		
		err = c.ShouldBindJSON(&transactionRecived)

		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 

		invalidParams := InValidParams(transactionRecived)

		// validamos que esten todos los parametros en el request
		if(len(invalidParams) > 0){
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,fmt.Sprintf("Faltan los campos %v", invalidParams)))
			return
		}



		tranUpdate,err := tran.service.Update(idTransaction,transactionRecived.Codigo,transactionRecived.Moneda,transactionRecived.Monto,
							transactionRecived.Emisor,transactionRecived.Receptor,transactionRecived.Fecha)

		
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
			return
		} 

		c.JSON(http.StatusOK, web.NewResponse(http.StatusOK,tranUpdate,""))

	}
}


func (tran Transaction) UpdateCodigoAndMonto() gin.HandlerFunc{
	return func(c *gin.Context) {

			err := ValidateToken(c.GetHeader("token"))
			if  err != nil {
				c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized,nil,err.Error()))
				return
			}
			idTransaction,err := strconv.Atoi(c.Param("id"))
	
			if err != nil {
				c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
				return
			} 
	
			var transactionRecived request
			
			err = c.ShouldBindJSON(&transactionRecived)
	
			if err != nil {
				c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
				return
			} 

			parametrosRequired := []string{"Codigo","Monto"}
			differences	:= ValidParms(transactionRecived,parametrosRequired)

			// validamos que esten todos los parametrosRequired en el body de la petecion
			if(len(differences) > 0){
				c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,fmt.Sprintf("Faltan los campos %v", differences)))
				return
			}

			
			tranUpdate,err := tran.service.UpdateCodigoAndMonto(idTransaction,transactionRecived.Codigo,
																transactionRecived.Monto)
	
			
			if err != nil {
				c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
				return
			} 
	
		
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK,tranUpdate,""))
	
		}
}

func (tran Transaction) Delete() gin.HandlerFunc{
	return func(c *gin.Context) {


		err := ValidateToken(c.GetHeader("token"))
		if  err != nil {
			c.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized,nil,err.Error()))

			return
		}

			idTransaction,err := strconv.Atoi(c.Param("id"))
	
			if err != nil {
				c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
				return
			} 

			err = tran.service.Delete(idTransaction)

			if err != nil {
				c.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest,nil,err.Error()))
				return
			} 
			c.JSON(http.StatusOK, web.NewResponse(http.StatusOK,fmt.Sprintf("Se elimino correctamente la transaccion con id: %v",idTransaction),""))
	}
}



func ValidParms(transactionRecived request,parametrosRequired []string) []string{
	var diff []string

	parametrosFromBody := GetParamsFromBody(transactionRecived)

        for _, s1 := range parametrosRequired {
            found := false
            for _, s2 := range parametrosFromBody {
                if s1 == s2 {
                    found = true
                    break
                }
            }
           
            if !found {
                diff = append(diff, s1)
            }
        }
    
        
    

    return diff
}



func GetParamsFromBody(parametros request) []string{
	var list []string
	r := reflect.ValueOf(parametros)

	for i := 0; i < r.NumField(); i++ {
		varValor := r.Field(i).Interface()
		if(varValor != "" && varValor != 0){

			list = append(list, r.Type().Field(i).Name)
		}

	}
	return list;
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

func filtrar(sliceTransaccion[]tra.Transaction, campo string, valor string) []tra.Transaction {
	var filtrado []tra.Transaction

	var per tra.Transaction
	fmt.Println(per)
	tipos := reflect.TypeOf(per)
	fmt.Println(tipos)
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


func ValidateToken(token string) error{

	if token == "" {
		return errors.New("token Vacio")
	}

	tokenFromEnv := os.Getenv("TOKEN")
	if token != tokenFromEnv {
		return errors.New("no tiene permisos para realizar la petición solicitada")
	}
	return nil
}
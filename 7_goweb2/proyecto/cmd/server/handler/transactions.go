package handler

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	internal "github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/internal/transactions"
	"github.com/gin-gonic/gin"
)

type request struct {
	CodigoDeTransaccion string  `json:"codigo_de_transaccion"`
	Moneda              string  `json:"moneda"`
	Monto               float64 `json:"monto"`
	Emisor              string  `json:"emisor"`
	Receptor            string  `json:"receptor"`
	FechaDeTransaccion  string  `json:"fecha_de_transaccion"`
}

type Transaction struct {
	service internal.Service
}

func NewTransaction(serv internal.Service) *Transaction {
	return &Transaction{service: serv}
}

func (t *Transaction) GetAll() gin.HandlerFunc { //TODO: implement filters
	return func(ctx *gin.Context) {
		response, err := t.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if len(response) == 0 {
			ctx.JSON(http.StatusOK, gin.H{})
			return
		}
		ctx.JSON(http.StatusOK, response)

	}
}

func (t *Transaction) GetTransactionByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid id",
			})
			return
		}
		response, err := t.service.GetTransactionByID(id)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, response)
		return

	}
}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var tr request
		err := ctx.ShouldBindJSON(&tr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid arguments",
			})
			return
		}
		err = validarCampos(tr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		respuesta, err := t.service.Store(tr.CodigoDeTransaccion, tr.Moneda, tr.Monto, tr.Emisor, tr.Receptor, tr.FechaDeTransaccion)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, respuesta)
		return

	}
}

func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id is mandatory",
			})
			return
		}

		var tr request
		err = ctx.ShouldBindJSON(&tr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid arguments",
			})
			return
		}
		err = validarCampos(tr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		respuesta, err := t.service.Update(id, tr.CodigoDeTransaccion, tr.Moneda, tr.Monto, tr.Emisor, tr.Receptor, tr.FechaDeTransaccion)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, respuesta)
		return

	}
}

func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id is mandatory",
			})
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.String(http.StatusOK, "transaction %d deleted", id)
		return

	}
}

func (t *Transaction) UpdateCodigoYMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "id is mandatory",
			})
			return
		}

		var tr request
		err = ctx.ShouldBindJSON(&tr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid arguments",
			})
			return
		}
		err = validarCampos(tr, "CodigoDeTransaccion", "Monto")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		respuesta, err := t.service.UpdateCodigoYMonto(id, tr.CodigoDeTransaccion, tr.Monto)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, respuesta)
		return

	}
}

//Funciones auxiliares

func (t *Transaction) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "" {
			if token == os.Getenv("TOKEN") { //"123456" {
				return
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Token incorrecto")

			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "No se ingresó un token")
	}
}

/*si campos es vacío, valida todos*/
func validarCampos(tr request, campos ...string) error {
	if len(campos) == 0 {
		campos = append(campos, "CodigoDeTransaccion", "Moneda", "Monto", "Emisor", "Receptor", "FechaDeTransaccion")
	}
	for _, campo := range campos {
		err := validarCampo(tr, campo)
		if err != nil {
			return err
		}
	}
	return nil
}

/*si campos es nil, valida todos*/
func validarCampo(tr request, campo string) error {
	switch campo {
	case "CodigoDeTransaccion":
		if tr.CodigoDeTransaccion == "" {
			return errors.New("el código de transacción es obligatorio")
		}
	case "Moneda":
		if tr.Moneda == "" {
			return errors.New("la moneda es obligatoria")
		}
	case "Monto":
		if tr.Monto == 0 {
			return errors.New("el monto es obligatorio")
		}
	case "Emisor":
		if tr.Emisor == "" {
			return errors.New("el emisor es obligatorio")
		}
	case "Receptor":
		if tr.Receptor == "" {
			return errors.New("el receptor es obligatorio")
		}
	case "FechaDeTransaccion":
		if tr.FechaDeTransaccion == "" {
			return errors.New("la fecha de transacción es obligatoria")
		}
	}
	return nil
}

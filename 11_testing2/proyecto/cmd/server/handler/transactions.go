package handler

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	internal "github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/internal/transactions"
	"github.com/extmatperez/meli_bootcamp2/11_testing2/proyecto/pkg/web"
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

// ListProducts godoc
// @Summary List transactions
// @Tags Transactions
// @Description get all transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response
// @Router /transactions [get]
func (t *Transaction) GetAll() gin.HandlerFunc { //TODO: implement filters
	return func(ctx *gin.Context) {

		filters := getFilters(ctx)

		response, err := t.service.GetAll(filters)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))

			return
		}
		ctx.JSON(200, web.NewResponse(200, response, ""))

	}
}

// ListProducts godoc
// @Summary Get transaction by ID
// @Tags Transactions
// @Description get transaction by ID
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "invalid id"
// @Failure 404 {object} web.Response "id not found"
// @Router /transactions:id [get]
func (t *Transaction) GetTransactionByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
			return
		}
		response, err := t.service.GetTransactionByID(id)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, response, ""))

	}
}

// StoreProducts godoc
// @Summary Store transaction
// @Tags Transactions
// @Description store transaction
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "invalid params"
// @Router /transactions [post]
func (t *Transaction) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var tr request
		err := ctx.ShouldBindJSON(&tr)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid arguments"))
			return
		}
		err = validarCampos(tr)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		respuesta, err := t.service.Store(tr.CodigoDeTransaccion, tr.Moneda, tr.Monto, tr.Emisor, tr.Receptor, tr.FechaDeTransaccion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, respuesta, ""))

	}
}

// ListProducts godoc
// @Summary Update transaction
// @Tags Transactions
// @Description update transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "invalid arguments"
// @Failure 404 {object} web.Response "transaction not found"
// @Router /transactions:id [put]
func (t *Transaction) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "El ID es obligatorio"))
			return
		}

		var tr request
		err = ctx.ShouldBindJSON(&tr)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid arguments"))
			return
		}
		err = validarCampos(tr)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		respuesta, err := t.service.Update(id, tr.CodigoDeTransaccion, tr.Moneda, tr.Monto, tr.Emisor, tr.Receptor, tr.FechaDeTransaccion)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, respuesta, ""))

	}
}

// ListProducts godoc
// @Summary Delete transaction
// @Tags Transactions
// @Description delete transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "invalid arguments"
// @Failure 404 {object} web.Response "transaction not found"
// @Router /transactions [put]
func (t *Transaction) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "El ID es obligatorio"))
			return
		}

		err = t.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, fmt.Sprintf("transaction %d deleted", id), ""))

	}
}

// ListProducts godoc
// @Summary Update Code and Amount of transaction
// @Tags Transactions
// @Description update transactions
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "invalid arguments"
// @Failure 404 {object} web.Response "transaction not found"
// @Router /transactions:id [patch]
func (t *Transaction) UpdateCodigoYMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "El ID es obligatorio"))

			return
		}

		var tr request
		err = ctx.ShouldBindJSON(&tr)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "Invalid arguments"))
			return
		}
		err = validarCampos(tr, "CodigoDeTransaccion", "Monto")
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		respuesta, err := t.service.UpdateCodigoYMonto(id, tr.CodigoDeTransaccion, tr.Monto)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, respuesta, ""))

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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token incorrecto"))

			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(401, nil, "No se ingresó un token"))
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

func getFilters(ctx *gin.Context) map[string]string {
	result := make(map[string]string)
	filter := ctx.Query("codigo")
	if filter != "" {
		result["Codigo"] = filter
	}
	filter = ctx.Query("moneda")
	if filter != "" {
		result["Moneda"] = filter
	}
	filter = ctx.Query("monto")
	if filter != "" {
		if _, err := strconv.ParseFloat(filter, 64); err != nil {
		} else {
			result["Monto"] = filter
		}
	}
	filter = ctx.Query("emisor")
	if filter != "" {
		result["Emisor"] = filter
	}
	filter = ctx.Query("receptor")
	if filter != "" {
		result["Receptor"] = filter
	}
	filter = ctx.Query("fecha_desde")
	if filter != "" {
		result["Fecha_desde"] = filter
	}
	filter = ctx.Query("fecha_hasta")
	if filter != "" {
		result["Fecha_hasta"] = filter
	}
	return result
}

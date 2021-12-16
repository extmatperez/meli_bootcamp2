package handler

import (
	"fmt"
	"os"
	"strconv"

	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/11_testing2/PracticaTT/internal/transacciones"
	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/9_goweb4/PracticaTT/pkg/web"
	"github.com/gin-gonic/gin"
)

type request struct {
	CodTransaccion string  `json:"cod_transaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         string  `json:"emisor"`
	Receptor       string  `json:"receptor"`
	FechaTrans     string  `json:"fecha_trans"`
}

type Transaccion struct {
	service transacciones.Service
}

func NewTransaccion(ser transacciones.Service) *Transaccion {
	return &Transaccion{service: ser}
}

// ListTransactions godoc
// @Summary List transactions
// @Tags Transactions
// @Description get transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions [get]
func (trans *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trans, err := trans.service.GetAll()
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewResponse(200, trans, ""))
		}
	}
}

// StoreTransactions godoc
// @Summary Store transactions
// @Tags Transactions
// @Description store transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions/load [post]
func (trans *Transaccion) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var transac request
		//valido el token
		// if !validarToken(ctx) {
		// 	return
		// }

		err := ctx.ShouldBindJSON(&transac)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error generando la persona. %v", err)))
		} else {
			trans, err := trans.service.Store(transac.CodTransaccion, transac.Moneda, transac.Monto, transac.Emisor, transac.Receptor, transac.FechaTrans)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error cargando la persona. %v", err)))
			} else {
				ctx.JSON(200, web.NewResponse(200, trans, ""))
			}
		}
	}
}

// SearchTransactions godoc
// @Summary Search transactions
// @Tags Transactions
// @Description search a transaction by id
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /transactions/:id [get]
func (trans *Transaccion) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		transac, err := trans.service.Search(param)

		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewResponse(200, transac, ""))
		}
	}
}

// FilterTransactions godoc
// @Summary Filter transactions
// @Tags Transactions
// @Description filter transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param params string
// @Success 200 {object} web.Response
// @Router /transactions/filter [get]
func (trans *Transaccion) Filter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Almaceno todas las etiquetas de mi struct
		var etiquetas []string
		etiquetas = append(etiquetas, "id", "cod_transaccion", "moneda", "monto", "emisor", "receptor", "fecha_trans")
		mapEtiquetas := make(map[string]string)
		mapRelacionEtiquetas := map[string]string{
			"id":              "Id",
			"cod_transaccion": "CodTransaccion",
			"moneda":          "Moneda",
			"monto":           "Monto",
			"emisor":          "Emisor",
			"receptor":        "Receptor",
			"fecha_trans":     "FechaTrans",
		}

		for _, etiqueta := range etiquetas {
			valEtiqueta := ctx.Query(etiqueta)
			if valEtiqueta != "" {
				mapEtiquetas[etiqueta] = valEtiqueta
			}
		}
		filtredTransaction, err := trans.service.Filter(mapEtiquetas, mapRelacionEtiquetas)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewResponse(200, filtredTransaction, ""))
		}
	}
}

// UpdateTransactions godoc
// @Summary Update transactions
// @Tags Transactions
// @Description update transactions
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Transaction to store"
// @Success 200 {object} web.Response
// @Router /transactions/:id [put]
func (trans *Transaccion) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//valido el token
		// if !validarToken(ctx) {
		// 	return
		// }

		//obtengo el id que quiero actualizar, y los datos a cambiar
		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error con el id %v:%v", param, err.Error())))
			return
		}
		var transac request
		err = ctx.ShouldBindJSON(&transac)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		transacResult, err := trans.service.Update(id, transac.CodTransaccion, transac.Moneda, transac.Monto, transac.Emisor, transac.Receptor, transac.FechaTrans)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, transacResult, ""))
	}
}

func (trans *Transaccion) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//valido el token
		// if !validarToken(ctx) {
		// 	return
		// }

		//obtengo el ID que se quiere eliminar
		paramId := ctx.Param("id")
		id, err := strconv.Atoi(paramId)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error con el id %v:%v", paramId, err.Error())))
			return
		}

		transacEliminated, err := trans.service.Delete(id)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, transacEliminated, ""))
		// ctx.String(200, "Se eliminó la transacción:")
		// ctx.JSON(200, transacEliminated)
	}
}

func (trans *Transaccion) UpdateCodigoYMonto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//valido el token
		if !validarToken(ctx) {
			return
		}

		//obtengo el id que quiero actualizar, y los datos a cambiar
		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error con el id %v:%v", param, err.Error())))
			return
		}
		var transac request
		err = ctx.ShouldBindJSON(&transac)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}

		//Valido los campos
		if transac.CodTransaccion == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El codigo de transacción no puede ser vacío"))
			return
		}
		if transac.Monto == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El monto no puede ser cero"))
			return
		}
		transacResult, err := trans.service.UpdateCodigoYMonto(id, transac.CodTransaccion, transac.Monto)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, transacResult)
	}
}

func validarToken(ctx *gin.Context) bool {
	token := ctx.Request.Header.Get("token")

	if token == "" {
		ctx.JSON(400, web.NewResponse(400, nil, "no se ha enviado ningun token"))
		return false
	}

	if token != os.Getenv("TOKEN") {
		ctx.JSON(404, web.NewResponse(404, nil, "token invalido"))
		return false
	}

	return true
}

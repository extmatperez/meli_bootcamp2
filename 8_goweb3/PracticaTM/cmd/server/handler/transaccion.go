package handler

import (
	"strconv"

	transacciones "github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/8_goweb3/PracticaTM/internal/transacciones"
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

func (trans *Transaccion) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trans, err := trans.service.GetAll()
		if err != nil {
			ctx.String(400, "Error: %v", err)
		} else {
			ctx.JSON(200, trans)
		}
	}
}

func (trans *Transaccion) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var transac request

		err := ctx.ShouldBindJSON(&transac)
		if err != nil {
			ctx.String(400, "Error generando la persona. %v", err)
		} else {
			trans, err := trans.service.Store(transac.CodTransaccion, transac.Moneda, transac.Monto, transac.Emisor, transac.Receptor, transac.FechaTrans)
			if err != nil {
				ctx.String(400, "Error cargando la persona. %v", err)
			} else {
				ctx.JSON(200, trans)
			}
		}
	}
}

func (trans *Transaccion) Search() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		param := ctx.Param("id")
		transac, err := trans.service.Search(param)

		if err != nil {
			ctx.String(404, "Error: %v", err.Error())
		} else {
			ctx.JSON(200, transac)
		}
	}
}

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
			ctx.String(200, err.Error())
		} else {
			ctx.JSON(200, filtredTransaction)
		}
	}
}

func (trans *Transaccion) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//valido el token
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.String(400, "no se ha enviado ningun token")
			return
		}
		if token != "aaa111" {
			ctx.String(400, "token invalido")
			return
		}

		//obtengo el id que quiero actualizar, y los datos a cambiar
		param := ctx.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			ctx.String(400, "Error con el id %v:%v", param, err.Error())
			return
		}
		var transac request
		err = ctx.ShouldBindJSON(&transac)
		if err != nil {
			ctx.String(400, err.Error())
			return
		}

		transacResult, err := trans.service.Update(id, transac.CodTransaccion, transac.Moneda, transac.Monto, transac.Emisor, transac.Receptor, transac.FechaTrans)
		if err != nil {
			ctx.JSON(404, gin.H{
				"Error": err.Error(),
			})
			return
		}
		ctx.JSON(200, transacResult)
	}
}

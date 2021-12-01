package handler
import (
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

func NewTransaction
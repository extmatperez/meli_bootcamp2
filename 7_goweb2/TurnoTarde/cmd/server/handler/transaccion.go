package handler
import (
	tran "https://github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/7_goweb2/TurnoTarde/internal/transaccion"
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

type Transaccion struct {
	service tran.Service
}

func NewPersona(ser personas.Service) *Persona {
	return &Persona{service: ser}
}

func (per *Persona) GetAll() gin.HandlerFunc {
	
}

func (controller *Persona) Store() gin.HandlerFunc {

}
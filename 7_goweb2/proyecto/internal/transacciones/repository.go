package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaCreacion     string  `json:"fecha_creacion"`
}

var transacciones []Transaccion

//var lastId int

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error)
	Load() ([]Transaccion, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// MIRAR ESTO MAÑANA
func (r *repository) Load() ([]Transaccion, error) {
	datos, err := os.ReadFile("./transaccion.json")
	json.Unmarshal(datos, &transacciones)
	return transacciones, err

}

func (r *repository) GetAll() ([]Transaccion, error) {
	return transacciones, nil
}

func (r *repository) Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error) {
	currentTime := time.Now().Format("02-01-2006")
	fecha := fmt.Sprintf(currentTime)
	codigo := uuid.NewV4().String()
	trans := Transaccion{id, codigo, moneda, monto, emisor, receptor, fecha}
	transacciones = append(transacciones, trans)
	return trans, nil
}

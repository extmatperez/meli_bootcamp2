package internal

import (
	"encoding/json"
	"os"
)

type Transaccion struct {
	ID                int     `json:"id"`
	CodigoTransaccion int     `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

var ts []Transaccion
var trID int

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string) (Transaccion, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaccion, error) {
	data, _ := os.ReadFile("6_goweb1/transacciones.json")
	json.Unmarshal(data, &ts)
	return ts, nil
}

func (r *repository) LastID() (int, error) {
	trID = len(ts) + 1
	return trID, nil
}

func (r *repository) Store(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string) (Transaccion, error) {
	t := Transaccion{id, codigo, moneda, monto, emisor, receptor, fecha}
	ts = append(ts, t)
	trID = t.ID

	return t, nil
}

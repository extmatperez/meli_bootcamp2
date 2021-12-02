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

type Repository interface {
	Load() ([]Transaccion, error)
	GetAll() ([]Transaccion, error)
	Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error)
	FindById(id int) (Transaccion, error)
	Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error)
	UpdateCod(id int, codigotransaccion string) (Transaccion, error)
	UpdateMon(id int, monto float64) (Transaccion, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Load() ([]Transaccion, error) {
	datos, err := os.ReadFile("../../internal/transacciones/transaccion.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(datos, &transacciones)
	if err != nil {
		return nil, err
	}
	return transacciones, nil

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

func (r *repository) FindById(id int) (Transaccion, error) {
	for i, v := range transacciones {
		if v.ID == id {
			return transacciones[i], nil
		}
	}
	return Transaccion{}, fmt.Errorf("La persona %d no existe", id)
}

func (r *repository) Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (Transaccion, error) {
	currentTime := time.Now().Format("02-01-2006")
	fecha := fmt.Sprintf(currentTime)
	codigo := uuid.NewV4().String()
	trans := Transaccion{id, codigo, moneda, monto, emisor, receptor, fecha}

	for i, v := range transacciones {
		if v.ID == id {
			transacciones[i] = trans
			return trans, nil
		}
	}
	return trans, fmt.Errorf("La persona %d no existe", id)
}

func (r *repository) UpdateCod(id int, codigotransaccion string) (Transaccion, error) {
	for i, v := range transacciones {
		if v.ID == id {
			transacciones[i].CodigoTransaccion = codigotransaccion
			return transacciones[i], nil
		}
	}
	return Transaccion{}, fmt.Errorf("La persona %d no existe", id)
}

func (r *repository) UpdateMon(id int, monto float64) (Transaccion, error) {
	for i, v := range transacciones {
		if v.ID == id {
			transacciones[i].Monto = monto
			return transacciones[i], nil
		}
	}
	return Transaccion{}, fmt.Errorf("La persona %d no existe", id)
}

func (r *repository) Delete(id int) error {
	for i, v := range transacciones {
		if v.ID == id {
			transacciones = append(transacciones[:i], transacciones[i+1:]...)
			return fmt.Errorf("La transaccion %d fue eliminada", id)
		}
	}
	return fmt.Errorf("la transaccion %d no existe", id)
}

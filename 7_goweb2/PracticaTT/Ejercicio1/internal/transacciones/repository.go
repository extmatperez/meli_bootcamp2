package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Transaccion struct {
	Id             int     `json:"id"`
	CodTransaccion string  `json:"cod_transaccion"`
	Moneda         string  `json:"moneda"`
	Monto          float64 `json:"monto"`
	Emisor         string  `json:"emisor"`
	Receptor       string  `json:"receptor"`
	FechaTrans     string  `json:"fecha_trans"`
}

var transacciones []Transaccion
var lastID int

type Repository interface {
	getAll() ([]Transaccion, error)
	Store(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
	LastId() (int, error)
	Search(id string) (Transaccion, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) getAll() ([]Transaccion, error) {
	return transacciones, nil
}

func (repo *repository) Store(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error) {
	trans := Transaccion{id, codTransaccion, moneda, monto, emisor, receptor, fechaTrans}
	err := repo.verificarCampos(trans)
	if err != nil {
		return Transaccion{}, err
	} else {
		lastID = id
		transacciones = append(transacciones, trans)
		return trans, nil
	}
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

func (repo *repository) verificarCampos(transac Transaccion) error {
	cadError := ""

	var campos []string
	campos = append(campos, "CodTransaccion", "Moneda", "Monto", "Emisor", "Receptor", "FechaTrans")

	//Recorro todos los campos definidos arriba de la transacción y genero error si está vacio alguno
	for _, campo := range campos {
		valor := reflect.ValueOf(transac).FieldByName(campo).Interface()
		if valor == "" {
			cadError += fmt.Sprintf("El campo %s es requerido\n", campo)
		}
	}

	if cadError != "" {
		return errors.New(cadError)
	}
	return nil
}

func (repo *repository) Search(id string) (Transaccion, error) {
	found := false
	for _, value := range transacciones {
		if strconv.Itoa(value.Id) == id {
			transac = value
			found = true
			break
		}
	}

	if found {
		return transac, nil
	} else {
		return transac, fmt.Errorf("no existe la transacción con el id %v", id)
	}
}

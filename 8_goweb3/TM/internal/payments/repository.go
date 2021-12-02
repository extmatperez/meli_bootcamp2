package internal

import (
	"fmt"
)

type Payment struct {
	Id       int     `json:"id"`
	Codigo   string  `json:"codigo"`
	Moneda   string  `json:"moneda"`
	Monto    float64 `json:"monto"`
	Emisor   string  `json:"emisor"`
	Receptor string  `json:"receptor"`
	Fecha    string  `json:"fecha"`
}

var payments []Payment
var lastId int

type Repository interface {
	GetAll() ([]Payment, error)
	Filter(codigo string, moneda string, monto float64, emisor string, receptor string, fecha string) ([]Payment, error)
	Store(id int, codigo string, moneda string, monto float64, emisor string, receptor string, fecha string) (Payment, error)
	Update(id int, codigo string, moneda string, monto float64, emisor string, receptor string, fecha string) (Payment, error)
	UpdateCodigo(id int, codigo string) (Payment, error)
	UpdateMonto(in int, monto float64) (Payment, error)
	Delete(id int) (string, error)
	LastId() (int, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Payment, error) {
	return payments, nil
}

func (repo *repository) Filter(codigo string, moneda string, monto float64, emisor string, receptor string, fecha string) ([]Payment, error) {
	filteredPayments := []Payment{}
	isFiltered := true
	for i, v := range payments {
		if codigo != "" && v.Codigo != codigo {
			isFiltered = false
		}
		if moneda != "" && v.Moneda != moneda {
			isFiltered = false
		}
		if monto != 0.0 && v.Monto != monto {
			isFiltered = false
		}
		if emisor != "" && v.Emisor != emisor {
			isFiltered = false
		}
		if receptor != "" && v.Receptor != receptor {
			isFiltered = false
		}
		if fecha != "" && v.Fecha != fecha {
			isFiltered = false
		}
		if isFiltered {
			filteredPayments = append(filteredPayments, payments[i])
		}
	}
	return filteredPayments, nil
}

func (repo *repository) Store(id int, codigo string, moneda string, monto float64, emisor string, receptor string, fecha string) (Payment, error) {
	pay := Payment{id, codigo, moneda, monto, emisor, receptor, fecha}
	lastId = id
	payments = append(payments, pay)
	return pay, nil
}

func (repo *repository) LastId() (int, error) {
	return lastId, nil
}

func (repo *repository) Update(id int, codigo string, moneda string, monto float64, emisor string, receptor string, fecha string) (Payment, error) {
	pay := Payment{id, codigo, moneda, monto, emisor, receptor, fecha}
	for i, v := range payments {
		if v.Id == id {
			payments[i] = pay
			return pay, nil
		}
	}
	return Payment{}, fmt.Errorf("La transacción %d no existe.", id)
}

func (repo *repository) UpdateCodigo(id int, codigo string) (Payment, error) {
	for i, v := range payments {
		if v.Id == id {
			payments[i].Codigo = codigo
			return payments[i], nil
		}
	}
	return Payment{}, fmt.Errorf("La transacción %d no existe.", id)
}

func (repo *repository) UpdateMonto(id int, monto float64) (Payment, error) {
	for i, v := range payments {
		if v.Id == id {
			payments[i].Monto = monto
			return payments[i], nil
		}
	}
	return Payment{}, fmt.Errorf("La transacción %d no existe.", id)
}

func (repo *repository) Delete(id int) (string, error) {
	index := 0
	for i, v := range payments {
		if v.Id == id {
			// Por defecto siempre comienza con 1 para que no se tenga que definir despues una suma en el slice.
			index = i
			// Y aca se sobreescribe con el contenido del slice que estaba antes del registro a borrar y todo lo que viene despues como un ellipsis.
			payments = append(payments[:index], payments[index+1:]...)
			return "Borrado correcto.", nil
		}
	}
	return "", fmt.Errorf("La transacción %d no existe.", id)
}

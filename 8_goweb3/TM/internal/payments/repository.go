package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/8_goweb3/TM/pkg/store"
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
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Payment, error) {
	err := repo.db.Read(&payments)

	if err != nil {
		return nil, err
	}

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
	repo.db.Read(&payments)

	pay := Payment{id, codigo, moneda, monto, emisor, receptor, fecha}

	payments = append(payments, pay)

	err := repo.db.Write(payments)

	if err != nil {
		return Payment{}, err
	}

	return pay, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&payments)

	if err != nil {
		return 0, nil
	}

	if len(payments) == 0 {
		return 0, nil
	}

	return payments[len(payments)-1].Id, nil
}

func (repo *repository) Update(id int, codigo string, moneda string, monto float64, emisor string, receptor string, fecha string) (Payment, error) {
	err := repo.db.Read(&payments)
	if err != nil {
		return Payment{}, err
	}

	pay := Payment{id, codigo, moneda, monto, emisor, receptor, fecha}
	for i, v := range payments {
		if v.Id == id {
			payments[i] = pay
			err := repo.db.Write(payments)
			if err != nil {
				return Payment{}, err
			}
			return pay, nil
		}
	}
	return Payment{}, fmt.Errorf("La transacción %d no existe.", id)
}

func (repo *repository) UpdateCodigo(id int, codigo string) (Payment, error) {
	err := repo.db.Read(&payments)
	if err != nil {
		return Payment{}, err
	}

	for i, v := range payments {
		if v.Id == id {
			payments[i].Codigo = codigo
			err := repo.db.Write(payments)
			if err != nil {
				return Payment{}, err
			}
			return payments[i], nil
		}
	}
	return Payment{}, fmt.Errorf("La transacción %d no existe.", id)
}

func (repo *repository) UpdateMonto(id int, monto float64) (Payment, error) {
	err := repo.db.Read(&payments)
	if err != nil {
		return Payment{}, err
	}
	for i, v := range payments {
		if v.Id == id {
			payments[i].Monto = monto
			err := repo.db.Write(payments)
			if err != nil {
				return Payment{}, err
			}
			return payments[i], nil
		}
	}
	return Payment{}, fmt.Errorf("La transacción %d no existe.", id)
}

func (repo *repository) Delete(id int) (string, error) {
	err := repo.db.Read(&payments)
	if err != nil {
		return "", err
	}

	index := 0
	for i, v := range payments {
		if v.Id == id {
			// Por defecto siempre comienza con 1 para que no se tenga que definir despues una suma en el slice.
			index = i
			// Y aca se sobreescribe con el contenido del slice que estaba antes del registro a borrar y todo lo que viene despues como un ellipsis.
			payments = append(payments[:index], payments[index+1:]...)
			err := repo.db.Write(payments)
			return "Borrado correcto.", nil
		}
	}
	return "", fmt.Errorf("La transacción %d no existe.", id)
}

package internal

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/11_testing2/TM/pkg/store"
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
	Filtrar(values ...string) ([]Payment, error)
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

			if err != nil {
				return "", err
			}

			return "Borrado correcto.", nil
		}
	}
	return "", fmt.Errorf("La transacción %d no existe.", id)
}

func (repo *repository) Filtrar(values ...string) ([]Payment, error) {
	var head []string
	head = append(head, "codigo", "moneda", "monto", "emisor", "receptor", "fecha")
	var filteredPayments []Payment

	filteredPayments = payments

	for i, v := range head {
		if len(values[i]) != 0 && len(filteredPayments) != 0 {
			filteredPayments = filtrarPayments(payments, v, values[i])
		}
		if len(filteredPayments) == 0 {
			return filteredPayments, fmt.Errorf("No hay coincidencias de transacciones con los filtros ingresados.")
		}
	}
	return filteredPayments, nil
}

func filtrarPayments(slicePayments []Payment, field string, value string) []Payment {
	var filteredPayments []Payment
	var pay Payment

	types := reflect.TypeOf(pay)
	i := 0
	for i = 0; i < types.NumField(); i++ {
		if strings.ToLower(types.Field(i).Name) == field {
			break
		}
	}
	for _, v := range slicePayments {
		var cadena string
		cadena = fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, value) {
			filteredPayments = append(filteredPayments, v)
		}

	}
	return filteredPayments
}

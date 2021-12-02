package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/extmatperez/meli_bootcamp2/tree/bouza_facundo/8_goweb3/PracticaTT/pkg/store"
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
	Filter(mapEtiquetas, mapRelacionEtiquetas map[string]string) ([]Transaccion, error)
	Update(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error)
	Delete(id int) (Transaccion, error)
	UpdateCodigoYMonto(id int, codTransaccion string, monto float64) (Transaccion, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) getAll() ([]Transaccion, error) {
	//Leo las transacciones desde store
	repo.db.Read(transacciones)
	return transacciones, nil
}

func (repo *repository) Store(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error) {
	//Leo las transacciones
	repo.db.Read(transacciones)

	trans := Transaccion{id, codTransaccion, moneda, monto, emisor, receptor, fechaTrans}
	err := repo.verificarCampos(trans)
	if err != nil {
		return Transaccion{}, err
	} else {
		lastID, err = repo.LastId()
		if err != nil {
			return Transaccion{}, err
		}
		transacciones = append(transacciones, trans)
		//Guardo las nuevas transacciones una vez agregada la nueva
		repo.db.Write(transacciones)

		return trans, nil
	}
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&transacciones)
	if err != nil {
		return 0, err
	}

	cantTrans := len(transacciones)
	if cantTrans == 0 {
		return 0, nil
	}

	lastID = transacciones[cantTrans-1].Id
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
	var transac Transaccion
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

func (repo *repository) Filter(mapEtiquetas, mapRelacionEtiquetas map[string]string) ([]Transaccion, error) {
	var filtredTransac []Transaccion
	var etiquetaStruct string
	for etiqueta, value := range mapEtiquetas {
		//Recorro cada etiqueta con su valor
		for _, transaccion := range transacciones {
			//Busco el valor de la etiqueta en las transacciones
			etiquetaStruct = mapRelacionEtiquetas[etiqueta]
			actValue := fmt.Sprintf("%v", reflect.ValueOf(transaccion).FieldByName(etiquetaStruct).Interface())
			if actValue == value {
				filtredTransac = append(filtredTransac, transaccion)
			}
		}
	}

	return filtredTransac, nil
}

func (repo *repository) Update(id int, codTransaccion, moneda string, monto float64, emisor, receptor, fechaTrans string) (Transaccion, error) {
	//Leo la transaccion desde el db
	err := repo.db.Read(&transacciones)
	if err != nil {
		return Transaccion{}, err
	}

	trans := Transaccion{id, codTransaccion, moneda, monto, emisor, receptor, fechaTrans}
	updated := false
	for i := range transacciones {
		if id == transacciones[i].Id {
			transacciones[i] = trans
			updated = true
			break
		}
	}
	if !updated {
		return Transaccion{}, fmt.Errorf("no se encontró la transaccion con el id %v", id)
	}

	//Una vez que modifico la transacción, lo guardo en la bd
	err = repo.db.Write(&transacciones)
	if err != nil {
		return Transaccion{}, err
	}

	return trans, nil
}

func (repo *repository) Delete(id int) (Transaccion, error) {
	var transEliminated Transaccion
	found := false
	for i, value := range transacciones {
		if value.Id == id {
			found = true
			transEliminated = value
			transacciones = append(transacciones[:i], transacciones[i+1:]...)
			break
		}
	}
	if !found {
		return Transaccion{}, fmt.Errorf("no se encontro la transacción con el id %v", id)
	}
	return transEliminated, nil
}

func (repo *repository) UpdateCodigoYMonto(id int, codTransaccion string, monto float64) (Transaccion, error) {
	var transacUpdated Transaccion
	found := false
	for i, _ := range transacciones {
		if transacciones[i].Id == id {
			found = true
			transacciones[i].CodTransaccion = codTransaccion
			transacciones[i].Monto = monto
			transacUpdated = transacciones[i]
		}
	}

	if !found {
		return Transaccion{}, fmt.Errorf("no se encontró la transaccion con el id %v", id)
	}
	return transacUpdated, nil
}

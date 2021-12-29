package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/internal/models"
	"github.com/extmatperez/meli_bootcamp2/17_storage1/proyecto/pkg/store"
	uuid "github.com/satori/go.uuid"
)

var transacciones []models.Transaccion

type Repository interface {
	Load() ([]models.Transaccion, error)
	GetAll() ([]models.Transaccion, error)
	Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error)
	FindById(id int) (models.Transaccion, error)
	FilterBy(valores ...string) ([]models.Transaccion, error)
	Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error)
	UpdateCod(id int, codigotransaccion string) (models.Transaccion, error)
	UpdateMon(id int, monto float64) (models.Transaccion, error)
	Delete(id int) error
	//DeleteAll() error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) Load() ([]models.Transaccion, error) {
	r.db.Read(&transacciones)
	datos, err := os.ReadFile("../../internal/transacciones/transaccion.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(datos, &transacciones)
	if err != nil {
		return nil, err
	}
	err = r.db.Write(transacciones)

	if err != nil {
		return nil, err
	}
	return transacciones, nil

}

func (r *repository) GetAll() ([]models.Transaccion, error) {
	err := r.db.Read(&transacciones)
	if err != nil {
		return nil, err
	}
	return transacciones, nil
}

func (r *repository) Store(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error) {
	r.db.Read(&transacciones)
	currentTime := time.Now().Format("02-01-2006")
	fecha := fmt.Sprint(currentTime)
	codigo := uuid.NewV4().String()
	trans := models.Transaccion{ID: id, CodigoTransaccion: codigo, Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor, FechaCreacion: fecha}
	transacciones = append(transacciones, trans)

	err := r.db.Write(transacciones)

	if err != nil {
		return models.Transaccion{}, err
	}

	return trans, nil
}

func (r *repository) FindById(id int) (models.Transaccion, error) {
	for i, v := range transacciones {
		if v.ID == id {
			return transacciones[i], nil
		}
	}
	return models.Transaccion{}, fmt.Errorf("la persona %d no existe", id)
}

func (r *repository) FilterBy(valores ...string) ([]models.Transaccion, error) {
	var etiquetas []string
	etiquetas = append(etiquetas, "moneda", "emisor", "receptor", "fechacreacion", "codigotransaccion")
	var transaccionesFiltradas []models.Transaccion

	transaccionesFiltradas = transacciones

	for i, v := range etiquetas {
		if len(valores[i]) != 0 && len(transaccionesFiltradas) != 0 {
			transaccionesFiltradas = filtrar(transacciones, v, valores[i])

		}
		if len(transaccionesFiltradas) == 0 {
			return transaccionesFiltradas, fmt.Errorf("no hay coincidencias")
		}
	}

	return transaccionesFiltradas, nil

}

func filtrar(sliceTransacciones []models.Transaccion, campo string, valor string) []models.Transaccion {
	var filtrado []models.Transaccion
	var trans models.Transaccion

	tipos := reflect.TypeOf(trans)
	i := 0
	for i = 0; i < tipos.NumField(); i++ {
		if strings.ToLower(tipos.Field(i).Name) == campo {
			break
		}
	}
	for _, v := range sliceTransacciones {
		var cadena string = fmt.Sprintf("%v", reflect.ValueOf(v).Field(i).Interface())
		if strings.Contains(cadena, valor) {
			filtrado = append(filtrado, v)
		}

	}
	return filtrado
}

func (r *repository) Update(id int, codigotransaccion string, moneda string, monto float64, emisor string, receptor string, fechacreacion string) (models.Transaccion, error) {
	err := r.db.Read(&transacciones)
	if err != nil {
		return models.Transaccion{}, err
	}
	currentTime := time.Now().Format("02-01-2006")
	fecha := fmt.Sprint(currentTime)
	codigo := uuid.NewV4().String()
	trans := models.Transaccion{ID: id, CodigoTransaccion: codigo, Moneda: moneda, Monto: monto, Emisor: emisor, Receptor: receptor, FechaCreacion: fecha}

	for i, v := range transacciones {
		if v.ID == id {
			transacciones[i] = trans
			r.db.Write(transacciones)
			return trans, nil
		}
	}
	return trans, fmt.Errorf("la persona %d no existe", id)
}

func (r *repository) UpdateCod(id int, codigotransaccion string) (models.Transaccion, error) {
	err := r.db.Read(&transacciones)
	if err != nil {
		return models.Transaccion{}, err
	}
	for i, v := range transacciones {
		if v.ID == id {
			transacciones[i].CodigoTransaccion = codigotransaccion
			r.db.Write(transacciones)
			return transacciones[i], nil
		}
	}
	return models.Transaccion{}, fmt.Errorf("la persona %d no existe", id)
}

func (r *repository) UpdateMon(id int, monto float64) (models.Transaccion, error) {
	err := r.db.Read(&transacciones)
	if err != nil {
		return models.Transaccion{}, err
	}
	for i, v := range transacciones {
		if v.ID == id {
			transacciones[i].Monto = monto
			r.db.Write(transacciones)
			return transacciones[i], nil
		}
	}
	return models.Transaccion{}, fmt.Errorf("la persona %d no existe", id)
}

func (r *repository) Delete(id int) error {
	err := r.db.Read(&transacciones)
	if err != nil {
		return err
	}
	for i, v := range transacciones {
		if v.ID == id {
			transacciones = append(transacciones[:i], transacciones[i+1:]...)
			r.db.Write(transacciones)

			return fmt.Errorf("la transaccion %d fue eliminada", id)
		}
	}
	return fmt.Errorf("la transaccion %d no existe", id)
}

/* func (r *repository) DeleteAll() error {
	err :=r.db.Read(&transacciones)
	 if err != nil {
		//return err
	}
	if len(transacciones) > 0 {
		transacciones = []Transaccion{}
		//r.db.Delete(transacciones)
		return fmt.Errorf("La transacciones fueron eliminadas")
	}
	return fmt.Errorf("No habia transacciones para eliminar")

} */

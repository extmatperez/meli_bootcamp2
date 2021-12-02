package internal

import (
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/7_goweb2/proyecto/pkg/store"
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

var trID int

type Repository interface {
	GetAll() ([]Transaccion, error)
	Store(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string) (Transaccion, error)
	LastID() (int, error)
	Update(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string)(Transaccion, error)
	UpdateEmisor(id int, emisor string)(Transaccion, error)
	Delete(id int) error
}

type repository struct{
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Transaccion, error) {
	var ts []Transaccion
	r.db.Read(&ts)
	return ts, nil
}

func (r *repository) LastID() (int, error) {
	var ts []Transaccion
	err := r.db.Read(&ts)
	if err != nil {
		return 0, err
	}

	return ts[len(ts)-1].ID, nil
}

func (r *repository) Store(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string) (Transaccion, error) {
	var ts []Transaccion
	err := r.db.Read(&ts)

	if err != nil {
		return Transaccion{}, err
	}

	t := Transaccion{id, codigo, moneda, monto, emisor, receptor, fecha}
	ts = append(ts, t)
	
	err = r.db.Write(ts)

	if err != nil {
		return Transaccion{}, err
	}

	return t, nil
}

func (r *repository) Update(id int, codigo int, moneda string, monto float64, emisor string, receptor string, fecha string) (Transaccion, error) {
	t := Transaccion{id, codigo, moneda, monto, emisor, receptor, fecha}
	var ts []Transaccion
	r.db.Read(&ts)
	for i := range ts{
		if ts[i].ID == id{
			t.ID = id
			ts[i] = t
			r.db.Write(&ts)
			return t, nil
		}
	}
	return Transaccion{}, fmt.Errorf("Producto %v no encontrado", id)
}

func(r *repository) UpdateEmisor(id int, emisor string)(Transaccion, error){
	var t Transaccion
	var ts []Transaccion
	r.db.Read(&ts)
	for i := range ts{
		if ts[i].ID == id{
			ts[i].Emisor = emisor
			t = ts[i]
			r.db.Write(&ts)
			return t, nil
		}
	}
	return Transaccion{}, fmt.Errorf("Producto %v no encontrado", id)
}

func(r *repository) Delete(id int) error{
	var index int
	var ts []Transaccion
	r.db.Read(&ts)
	for i := range ts{
		if ts[i].ID == id{
			index = i
			ts = append(ts[:index], ts[index+1:]...)
			r.db.Write(ts)
			return nil
		}
	}
	return fmt.Errorf("Producto %v no encontrado", id)
}

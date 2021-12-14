package internal

import (
	"errors"
	"fmt"

	"github.com/extmatperez/meli_bootcamp2/tree/de_bonis_matias/8_goweb3/TT/go-web/pkg/store"
)

type Producto struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Color     string `json:"color"`
	Precio    string `json:"precio"`
	Stock     int    `json:"stock"`
	Codigo    string `json:"codigo"`
	Publicado bool   `json:"publicado"`
	Creado    string `json:"creado"`
}

var ps []Producto

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
	Edit(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
	LastID() (int, error)
	Delete(id int) error
	Change(id int, nombre, precio string) (Producto, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Producto, error) {
	err := r.db.Read(&ps)
	if err != nil {
		return []Producto{}, err
	}
	return ps, nil
}

func (repo *repository) LastID() (int, error) {
	err := repo.db.Read(&ps)

	if err != nil {
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].ID + 1, nil
}

func (r *repository) Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error) {
	var ps []Producto
	err := r.db.Read(&ps)
	if err != nil {
		return Producto{}, err
	}
	newId, err := r.LastID()
	if err != nil {
		return Producto{}, err
	}
	fmt.Println(newId)
	p := Producto{newId, nombre, color, precio, stock, codigo, publicado, creado}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Producto{}, err
	}
	return p, nil
}

func (r *repository) Edit(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error) {
	pEdit := Producto{id, nombre, color, precio, stock, codigo, publicado, creado}
	err := r.db.Read(&ps)
	if err != nil {
		return Producto{}, err
	}
	for i, p := range ps {
		if p.ID == id {
			ps[i] = pEdit
			err := r.db.Write(ps)
			if err != nil {
				return Producto{}, err
			}
			return pEdit, nil
		}
	}
	errText := fmt.Sprintf("El usuario %d no existe", id)
	return Producto{}, errors.New(errText)
}

func (r *repository) Delete(id int) error {
	index := 0
	err := r.db.Read(&ps)
	if err != nil {
		return err
	}
	for i, v := range ps {
		if v.ID == id {
			index = i
			ps = append(ps[:index], ps[index+1:]...)
			err := r.db.Write(ps)
			return err
		}
	}
	return fmt.Errorf("la persona %d no existe", id)
}

func (e *repository) Change(id int, nombre, precio string) (Producto, error) {
	for i, p := range ps {
		if p.ID == id {
			if nombre != "" {
				ps[i].Nombre = nombre
			}
			if precio != "" {
				ps[i].Precio = precio
			}
			return ps[i], nil
		}
	}
	errText := fmt.Sprintf("El usuario %d no existe", id)
	return Producto{}, errors.New(errText)
}

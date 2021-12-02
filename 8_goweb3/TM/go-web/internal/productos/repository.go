package internal

import (
	"errors"
	"fmt"
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
var lastID int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
	Edit(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error)
	LastID() (int, error)
	Delete(id int) error
	Change(id int, nombre, precio string) (Producto, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Producto, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error) {
	p := Producto{id, nombre, color, precio, stock, codigo, publicado, creado}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) Edit(id int, nombre, color, precio string, stock int, codigo string, publicado bool, creado string) (Producto, error) {
	pEdit := Producto{id, nombre, color, precio, stock, codigo, publicado, creado}
	for i, p := range ps {
		if p.ID == id {
			ps[i] = pEdit
			return pEdit, nil
		}
	}
	errText := fmt.Sprintf("El usuario %d no existe", id)
	return Producto{}, errors.New(errText)
}

func (r *repository) Delete(id int) error {
	index := 0
	for i, v := range ps {
		if v.ID == id {
			index = i
			ps = append(ps[:index], ps[index+1:]...)
			return nil
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

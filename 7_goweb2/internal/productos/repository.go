package productos

import (
	"errors"
	"time"
)

type Producto struct {
	Id        int       `json:"id"`
	Nombre    string    `json:"nombre"`
	Color     string    `json:"color"`
	Precio    float64   `json:"precio"`
	Stock     int       `json:"stock"`
	Codigo    string    `json:"codigo"`
	Publicado bool      `json:"publicado"`
	Creado    time.Time `json:"created_at"`
}

type Repository interface {
	GetAll() ([]Producto, error)
	Get(id int) (Producto, error)
	Store(nombre, color string, precio float64, stock int, codigo string, publicado bool) (Producto, error)
}

type repository struct {
	Productos []Producto
	LastId    int
}

func NewRepository() Repository {
	return &repository{
		Productos: []Producto{
			{1, "Tijera Negra", "negro", 10.0, 3, "A15", true, time.Now()},
			{2, "Remera Blanca", "blanco", 10.0, 3, "A16", true, time.Now()},
		},
		LastId: 2,
	}
}

func (r *repository) GetAll() ([]Producto, error) {
	return r.Productos, nil
}

func (r *repository) Get(id int) (Producto, error) {
	for _, p := range r.Productos {
		if p.Id == id {
			return p, nil
		}
	}

	return Producto{}, errors.New("Not found")
}

func (r *repository) Store(
	nombre, color string,
	precio float64,
	stock int,
	codigo string,
	publicado bool) (Producto, error) {

	producto := Producto{
		Id:        r.LastId + 1,
		Nombre:    nombre,
		Color:     color,
		Precio:    precio,
		Stock:     stock,
		Codigo:    codigo,
		Publicado: publicado,
		Creado:    time.Now(),
	}

	r.Productos = append(r.Productos, producto)
	r.LastId += 1
	return producto, nil
}

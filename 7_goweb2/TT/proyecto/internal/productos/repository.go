/*
Repositorio, debe tener el acceso a la variable guardada en memoria.
a. Se debe crear el archivo repository.go
b. Se debe crear la estructura de la entidad
c. Se deben crear las variables globales donde guardar las entidades
d. Se debe generar la interface Repository con todos sus métodos
e. Se debe generar la estructura repository
f. Se debe generar una función que devuelva el Repositorio
g. Se deben implementar todos los métodos correspondientes a las operaciones
a realizar (GetAll, Store, etc..)
*/

package internal

import (
	"encoding/json"
	"os"
)

type Producto struct {
	ID            int    `json:"id"`
	Nombre        string `json:"nombre"`
	Color         string `json:"color"`
	Precio        string `json:"precio"`
	Stock         int    `json:"stock"`
	Codigo        string `json:"codigo"`
	Publicado     bool   `json:"publicado"`
	FechaCreacion string `json:"fechaCreacion"`
}

var productos []Producto
var lastId int

type Repository interface {
	GetAll() ([]Producto, error)
	Store(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	LastId() (int, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func obtenerProductos() []Producto {

	data, _ := os.ReadFile("7_goweb2/productos.json")

	var lista []Producto

	json.Unmarshal(data, &lista)

	return lista
}

func escribirJSON(nuevaLista []Producto) error {

	nuevaListaFormateada, _ := json.Marshal(nuevaLista)

	err := os.WriteFile("7_goweb2/productos.json", nuevaListaFormateada, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll() ([]Producto, error) {

	productos = obtenerProductos()

	return productos, nil
}

func (r *repository) Store(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	nuevoProducto := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	lastId = id

	productos = append(productos, nuevoProducto)

	err := escribirJSON(productos)

	if err != nil {
		return Producto{}, err
	}

	return nuevoProducto, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

/*
Se solicita implementar una funcionalidad que modifique completamente una entidad. Para
lograrlo, es necesario, seguir los siguientes pasos:
	1. Generar un método PUT para modificar la entidad completa
	2. Desde el Path enviar el ID de la entidad que se modificará
	3. En caso de no existir, retornar un error 404
	4. Realizar todas las validaciones (todos los campos son requeridos)
*/

/*
Es necesario implementar una funcionalidad para eliminar una entidad. Para lograrlo, es
necesario, seguir los siguientes pasos:
	1. Generar un método DELETE para eliminar la entidad en base al ID
	2. En caso de no existir, retornar un error 404
*/

/*
Se requiere implementar una funcionalidad que modifique la entidad parcialmente, solo se
deben modificar 2 campos:
	- Si se seleccionó Productos, los campos nombre y precio.
	- Si se seleccionó Usuarios, los campos apellido y edad.
	- Si se seleccionó Transacciones, los campos código de transacción y monto.

Para lograrlo, es necesario, seguir los siguientes pasos:
	1. Generar un método PATCH para modificar la entidad parcialmente, modificando solo 2
	campo (a elección)
	2. Desde el Path enviar el ID de la entidad que se modificara
	3. En caso de no existir, retornar un error 404
	4. Realizar las validaciones de los 2 campos a enviar
*/

package internal

import (
	"encoding/json"
	"fmt"
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
	Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error)
	Delete(id int) error
	UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func cargarProductos() {

	data, _ := os.ReadFile("/Users/beconti/Desktop/meli_bootcamp2/7_goweb2/productos.json")

	json.Unmarshal(data, &productos)

}

func escribirJSON(nuevaLista []Producto) error {

	nuevaListaFormateada, _ := json.Marshal(nuevaLista)

	err := os.WriteFile("/Users/beconti/Desktop/meli_bootcamp2/7_goweb2/productos.json", nuevaListaFormateada, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAll() ([]Producto, error) {

	cargarProductos()

	return productos, nil
}

func (r *repository) Store(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	cargarProductos()

	nuevoProducto := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	productos = append(productos, nuevoProducto)

	err := escribirJSON(productos)

	if err != nil {
		return Producto{}, err
	}

	return nuevoProducto, nil
}

func (r *repository) LastId() (int, error) {

	cargarProductos()

	lastId = productos[len(productos)-1].ID

	return lastId, nil
}

func (r *repository) Update(id int, nombre string, color string, precio string, stock int, codigo string, publicado bool, fechaCreacion string) (Producto, error) {

	cargarProductos()

	productoActualizado := Producto{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}

	for i, p := range productos {
		if p.ID == id {
			productos[i] = productoActualizado
			escribirJSON(productos)
			return productoActualizado, nil
		}
	}

	return Producto{}, fmt.Errorf("el producto %v no existe", id)

}

func (r *repository) Delete(id int) error {

	cargarProductos()

	index := 0

	for i, p := range productos {

		if p.ID == id {
			index = i
			productos = append(productos[:index], productos[index+1:]...)
			escribirJSON(productos)
			return nil
		}

	}

	return fmt.Errorf("el producto %v no existe", id)
}

func (r *repository) UpdateNombrePrecio(id int, nombre string, precio string) (Producto, error) {

	cargarProductos()

	for i, p := range productos {

		if p.ID == id {
			productos[i].Nombre = nombre
			productos[i].Precio = precio

			escribirJSON(productos)

			return productos[i], nil
		}
	}

	return Producto{}, fmt.Errorf("el producto %v no existe", id)
}

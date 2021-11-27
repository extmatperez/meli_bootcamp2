/*
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y
retornar el valor del precio total.

Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio

Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.

Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.

Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.

Interface Ecommerce:
- El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
- El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/

package main

import "fmt"

const (
	Pequenio = "pequenio"
	Mediano  = "mediano"
	Grande   = "grande"
)

type Tienda struct {
	productos []Producto
}

type Producto struct {
	tipoProducto string
	nombre       string
	precio       float64
}

type InterfazProducto interface {
	calcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func nuevoProducto(tipoProducto, nombre string, precio float64) Producto {
	return Producto{tipoProducto, nombre, precio}
}

func nuevaTienda(productos []Producto) Ecommerce {
	nuevaTienda := Tienda{productos}

	return &nuevaTienda
}

func (p Producto) calcularCosto() float64 {

	switch p.tipoProducto {
	case Pequenio:
		return p.precio
	case Mediano:
		return p.precio + p.precio*0.03
	case Grande:
		return p.precio + p.precio*0.06 + 2500.0
	default:
		return 0.0
	}

}

func (t Tienda) Total() float64 {

	total := 0.0

	for _, p := range t.productos {
		total += p.calcularCosto()
	}

	return total
}

func (t *Tienda) Agregar(p Producto) {
	t.productos = append(t.productos, p)
}

func main() {

	p1 := nuevoProducto("pequenio", "Laptop", 500)
	p2 := nuevoProducto("mediano", "Mouse", 150)
	p3 := nuevoProducto("grande", "Teclado", 250)

	fmt.Printf("Precio p1: $%f\n", p1.calcularCosto())
	fmt.Printf("Precio p2: $%f\n", p2.calcularCosto())
	fmt.Printf("Precio p3: $%f\n", p3.calcularCosto())

	nuevaTienda := nuevaTienda([]Producto{p1, p2})
	fmt.Printf("Tienda creada: %v\n", nuevaTienda)

	nuevaTienda.Agregar(p3)
	fmt.Printf("Agregamos un nuevo producto: %v\n", nuevaTienda)

	fmt.Printf("El costo total de la tienda es $%v\n", nuevaTienda.Total())
}

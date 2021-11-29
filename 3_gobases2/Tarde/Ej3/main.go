/*
? Ejercicio 3 - Productos
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar
productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos
de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío
de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y
devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y
los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda


*/

package main

import "fmt"

type tienda struct {
	ListaProductos []Producto
}

func (t *tienda) Total() (tota float64) {
	total := 0.0
	for _, prod := range t.ListaProductos {
		total += prod.CalcularCosto()
	}
	return
}

func (t *tienda) Agregar(p Producto) {
	t.ListaProductos = append(t.ListaProductos, p)
}

type tipo interface {
	CalcularAdicional() func(precio float64) float64
}

type grande struct{}
type mediano struct{}
type chico struct{}

func (t grande) CalcularAdicional() func(precio float64) float64 {
	funcion := func(precio float64) float64 {
		return precio*1.06 + 2500.0
	}
	return funcion
}

func (t mediano) CalcularAdicional() func(precio float64) float64 {
	return func(precio float64) float64 {
		return precio * 1.03
	}
}

func (t chico) CalcularAdicional() func(precio float64) float64 {
	return func(precio float64) float64 {
		return precio
	}
}

type producto struct {
	Nombre string
	Precio float64
	Tipo   tipo
}

func (p producto) CalcularCosto() float64 {
	return p.Tipo.CalcularAdicional()(p.Precio)
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func nuevoProducto(tipoProducto string, nombre string, precio float64) Producto {
	switch tipoProducto {
	case "Grande":
		// tip := grande{}
		// return producto{tip, nombre, precio}
	case "Mediano":
		// tip := mediano{}
		// return producto{tip, nombre, precio}
	case "Chico":
		// tip := chico{}
		// return producto{tip, nombre, precio}
	}
	return producto{}
}

func NuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	ecom := NuevaTienda()
	ecom.Agregar(nuevoProducto("Chico", "coca", 100.0))
	fmt.Printf("Tienda 1: Total: %.2f\n", ecom.Total())
	// ecom.Agregar(nuevoProducto("Mediano", "arroz", 100.0))
	// fmt.Printf("Tienda 2: Total: %.2f\n", ecom.Total())
	// ecom.Agregar(nuevoProducto("Grande", "leche", 100.0))
	// fmt.Printf("Tienda 2: Total: %.2f\n", ecom.Total())

	ecom2 := NuevaTienda()
	// ecom2.Agregar(nuevoProducto("Chico", "coca", 100.0))
	// fmt.Printf("Tienda 1: Total: %.2f\n", ecom2.Total())
	ecom2.Agregar(nuevoProducto("Mediano", "arroz", 100.0))
	fmt.Printf("Tienda 2: Total: %.2f\n", ecom2.Total())
	// ecom2.Agregar(nuevoProducto("Grande", "leche", 100.0))
	// fmt.Printf("Tienda 2: Total: %.2f\n", ecom2.Total())

	ecom3 := NuevaTienda()
	// ecom3.Agregar(nuevoProducto("Chico", "coca", 100.0))
	// fmt.Printf("Tienda 1: Total: %.2f\n", ecom3.Total())
	// ecom3.Agregar(nuevoProducto("Mediano", "arroz", 100.0))
	// fmt.Printf("Tienda 2: Total: %.2f\n", ecom3.Total())
	ecom3.Agregar(nuevoProducto("Grande", "leche", 100.0))
	fmt.Printf("Tienda 2: Total: %.2f\n", ecom3.Total())
}

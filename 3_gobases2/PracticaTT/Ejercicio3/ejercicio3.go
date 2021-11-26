package main

/*
Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.

Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/
type Producto interface {
	calcularCosto() float64
}

type Ecommerce interface {
	total()
	agregar()
}

type tienda struct {
	productos []Producto
}

func (t tienda) total() {

}

func (t tienda) agregar() {

}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

type pequeño struct {
	p producto
}

type mediano struct {
	p producto
}

type grande struct {
	p producto
}

func nuevoProducto(tipo, nombre string, precio float64) producto {
	prod := producto{tipo, nombre, precio}
	return prod
}

func nuevaTienda() Ecommerce {
	return tienda{}
}

func (p producto) calcularCosto() float64 {

}

func main() {

}

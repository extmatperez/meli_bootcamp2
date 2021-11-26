package main

import (
	"fmt"
)

type Producto struct {
	tipo   string
	name   string
	precio float64
}
type IProducto interface {
	CalculaCosto()
}
type Ecomerece interface {
	total()
	agregar()
}

type tienda struct {
	productos []Producto
}

func nuevoProducto(tipoProducto, nombre string, precio float64) Producto {
	nuevoProducto := Producto{tipoProducto, nombre, precio}

	return nuevoProducto
}

func nuevaTienda(productos ...Producto) tienda {
	t := tienda{productos}

	return t
}

func (producto *Producto) CalculaCosto() float64 {
	adicional := map[string]float64{
		"mouse":     producto.precio,
		"teclado":   producto.precio + (producto.precio * 0.03),
		"parlantes": producto.precio + (producto.precio * 0.06) + 2500,
	}

	return adicional[producto.tipo]
}

func (tienda tienda) Total() float64 {
	resultado := 0.0

	for _, p := range tienda.productos {
		resultado += p.precio
	}

	return resultado
}

func (tienda *tienda) Add(producto Producto) []Producto {
	tienda.productos = append(tienda.productos, producto)

	return tienda.productos
}

func main() {
	p1 := nuevoProducto("mouse", "mac", 33000)
	p2 := nuevoProducto("teclado", "gamer", 120000)
	p3 := nuevoProducto("parlantes", "parlante doble", 200000)

	fmt.Printf("Total Producto 1: %v\n", p1.CalculaCosto())
	fmt.Printf("Totla Producto 2: %v\n", p2.CalculaCosto())
	fmt.Printf("Total Producto 3 : %v\n", p3.CalculaCosto())

	t1 := nuevaTienda(p1, p2)
	fmt.Printf("tienda: %v\n", t1)

	t1.Add(p3)
	fmt.Printf("Producto nuevo en la tienda: %v\n", t1)

	fmt.Printf("Total: %v\n", t1.Total())

}

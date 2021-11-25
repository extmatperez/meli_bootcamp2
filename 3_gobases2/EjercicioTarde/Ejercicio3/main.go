package main

import "fmt"

const (
	Pequeno = "Pequeno"
	Mediano = "Mediano"
	Grande  = "Grande"
)

type tienda struct {
	productos []Producto
}

func (t *tienda) Agregar(p Producto) {
	t.productos = append(t.productos, p)
}

func (t tienda) Total() float64 {
	total := 0.0
	for _, valor := range t.productos {
		total += valor.CalcularCosto()
	}
	return total
}

type producto struct {
	Tipo   string
	Nombre string
	Precio int
}

type pequeno struct {
	P producto
}

type mediano struct {
	P producto
}

type grande struct {
	P producto
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func (p pequeno) CalcularCosto() float64 {
	return float64(p.P.Precio)
}

func (p mediano) CalcularCosto() float64 {
	return float64(p.P.Precio) + (float64(p.P.Precio) * 0.03)
}

func (p grande) CalcularCosto() float64 {
	return float64(p.P.Precio) + (float64(p.P.Precio) * 0.06) + 2500.0
}

func nuevoProducto(tipo string, nombre string, precio int) Producto {

	switch tipo {
	case Pequeno:
		p := pequeno{producto{Tipo: tipo, Nombre: nombre, Precio: precio}}
		return p
	case Mediano:
		m := mediano{producto{Tipo: tipo, Nombre: nombre, Precio: precio}}
		return m
	case Grande:
		g := grande{producto{Tipo: tipo, Nombre: nombre, Precio: precio}}
		return g
	}
	return nil
}

func totalTienda(e Ecommerce) {
	fmt.Printf("Total: %.2f\n\n", e.Total())
}

func agregarATienda(e Ecommerce, p Producto) {
	e.Agregar(p)
}

func main() {

	var tienda1 tienda

	agregarATienda(&tienda1, nuevoProducto(Pequeno, "Producto1", 1000))
	agregarATienda(&tienda1, nuevoProducto(Grande, "Producto2", 20000))
	agregarATienda(&tienda1, nuevoProducto(Mediano, "Producto3", 10000))
	agregarATienda(&tienda1, nuevoProducto(Pequeno, "Producto4", 500))
	fmt.Printf("\nTienda1: %+v \n", tienda1)
	totalTienda(&tienda1)

	var tienda2 tienda

	agregarATienda(&tienda2, nuevoProducto(Pequeno, "Producto5", 2000))
	agregarATienda(&tienda2, nuevoProducto(Grande, "Producto6", 40000))
	agregarATienda(&tienda2, nuevoProducto(Mediano, "Producto7", 15000))
	agregarATienda(&tienda2, nuevoProducto(Pequeno, "Producto8", 5000))
	fmt.Printf("Tienda2: %+v \n", tienda2)
	totalTienda(&tienda2)
}

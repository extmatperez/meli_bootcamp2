package main

import "fmt"

type tienda struct {
	ListaProductos []producto
}

func (t tienda) Total() float64 {
	montoTodal := 0.0
	for _, value := range t.ListaProductos {

		switch value.tipo {
		case "Pequeño":
			montoTodal = montoTodal + value.precio
			continue
		case "Mediano":
			montoTodal = montoTodal + value.precio + (value.precio * 0.03)
			continue
		case "Grande":
			montoTodal = montoTodal + value.precio + (value.precio * 0.06) + 2500
			continue
		}
	}
	return montoTodal
}
func (t tienda) Agregar(p producto) {
	t.ListaProductos = append(t.ListaProductos, p)
}

type producto struct {
	tipo   string
	nombre string
	precio float64
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case "Pequeño":
		return 0.0
	case "Mediano":
		return (p.precio * 0.03)
	case "Grande":
		return (p.precio * 0.06) + 2500

	}
	return 0.0
}

type Producto interface {
	CalcularCosto() float64
}
type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

func nuevoProducto(tipo, nombre string, precio float64) producto {
	p := producto{tipo, nombre, precio}
	return p
}
func nuevaTienda() Ecommerce {
	eco := tienda{}
	return eco
}

func main() {

	produc := nuevoProducto("Mediano", "Perfume", 200)

	fmt.Println(produc)
	fmt.Println(produc.CalcularCosto())

	tiendita := nuevaTienda()
	fmt.Println(tiendita)

	tiendita.Agregar(produc)
	fmt.Println(tiendita)

	fmt.Println(tiendita.Total())

}

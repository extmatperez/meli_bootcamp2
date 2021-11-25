package main

import "fmt"

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type tienda struct {
	Lista []producto
}

func (t *tienda) Total() float64 {
	total := 0.0
	for _, prod := range t.Lista {
		total += prod.CalcularCosto(prod)
	}
	return total
}

func (t *tienda) Agregar(prod producto) {
	list := append(t.Lista, prod)
	fmt.Println("Producto agregado", list)
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

func (p *producto) CalcularCosto(prod producto) float64 {
	switch prod.Tipo {
	case pequeño:
		return prod.Precio
	case mediano:
		return prod.Precio * 0.03
	case grande:
		return prod.Precio*0.06 + 2500
	default:
		return 0.0
	}
}

type Producto interface {
	CalcularCosto()
}

type Ecommerce interface {
	Total()
	Agregar()
}

func nuevoProducto(tipo string, nombre string, precio float64) producto {
	prod := producto{Tipo: tipo, Nombre: nombre, Precio: precio}
	return prod
}

func nuevaTienda() tienda {
	ecom := tienda{}
	return ecom
}

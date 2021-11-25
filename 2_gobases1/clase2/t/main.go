package main

import (
	"errors"
	"fmt"
)

// ej1
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      string
	Fecha    string
}

// ej3
const (
	small  = "pequeno"
	medium = "mediano"
	big    = "grande"
)

type tienda struct {
	Products []producto
}

func (t tienda) Total() float64 {
	total := 0.0
	for _, product := range t.Products {
		costo, err := product.CalcularCosto()
		if err == nil {
			total += costo
		}
	}

	return total
}

func (t *tienda) Agregar(p producto) {
	t.Products = append(t.Products, p)
}

type producto struct {
	Tipo   string
	Nombre string
	Precio float64
}

type Producto interface {
	CalcularCosto() float64
}

func (p producto) CalcularCosto() (float64, error) {
	if p.Tipo == small {
		return p.Precio, nil
	}
	if p.Tipo == medium {
		return p.Precio * 1.03, nil
	}
	if p.Tipo == big {
		return p.Precio*1.06 + 2500, nil
	}
	return 0, errors.New("tipo inexistente")
}

type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

func nuevoProducto(tipo string, nombre string, precio float64) producto {
	return producto{Tipo: tipo, Nombre: nombre, Precio: precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	tiendaUno := nuevaTienda()
	agua := nuevoProducto(small, "agua destilada", 100.00)
	aceite := nuevoProducto(medium, "aceite", 300.00)
	neumatico := nuevoProducto(big, "neumatico", 5000.00)
	tiendaUno.Agregar(agua)
	tiendaUno.Agregar(aceite)
	tiendaUno.Agregar(aceite)
	tiendaUno.Agregar(aceite)
	tiendaUno.Agregar(neumatico)

	total := tiendaUno.Total()

	fmt.Println(total)
}

package main

import (
	"errors"
	"fmt"
)

type producto struct {
	Nombre string
	Precio float64
	Tipo   string
}

type tienda struct {
	Productos []producto
}

type Producto interface {
	CalcularCosto() int
}

type Ecommerce interface {
	Total() float64
	Agregar()
}

func nuevoProducto(tipo, nombre string, precio float64) producto {
	//producto1 := producto{nombre, precio, tipo}
	producto2 := producto{
		Nombre: nombre,
		Precio: precio,
		Tipo:   tipo,
	}

	return producto2
}

func nuevaTienda() tienda {
	tienda1 := tienda{}
	return tienda1
}

func CalcularCosto(p producto) (float64, error) {
	switch p.Tipo {
	case "Pequeño":
		return p.Precio, nil
	case "Mediano":
		return p.Precio * 1.03, nil
	case "Grande":
		return p.Precio*1.06 + 2500, nil
	default:
		return 0.0, errors.New("tipo de producto no valido")
	}
}

func (t tienda) Total() float64 {
	total := 0.0
	for _, p := range t.Productos {
		res, err := CalcularCosto(p)
		if err != nil {
			fmt.Println(err)
		} else {
			total += res
		}
	}
	return total
}

func (t *tienda) Agregar(p producto) {
	t.Productos = append(t.Productos, p)
}

func main() {

	fmt.Println("Bienvenido al ejercicio 3")

}

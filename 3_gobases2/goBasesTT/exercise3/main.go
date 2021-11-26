package main

import "fmt"

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

func main() {
	var local = NuevaTienda()

	tipo := []string{pequeño, mediano, grande}

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			p := NuevoProducto(tipo[j], "Alfajor", float64(i*i))
			local.Agregar(p)

		}
	}

	fmt.Println(local.Total())
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

type tienda struct {
	Productos []Producto
}

type Producto interface {
	CalcularCosto() float64
}

type producto struct {
	tipo, nombre string
	precio       float64
}

func NuevoProducto(tipos, nombres string, precios float64) producto {
	prod := producto{tipo: tipos, nombre: nombres, precio: precios}
	return prod
}

func NuevaTienda() Ecommerce {
	var productos []Producto = make([]Producto, 0)
	var local = tienda{productos}
	var ecom Ecommerce = &local
	return ecom
}

func (t *tienda) Agregar(p Producto) {
	t.Productos = append(t.Productos, p)
}

func (t tienda) Total() float64 {
	var total float64 = 0.0

	for _, p := range t.Productos {
		total = total + p.CalcularCosto()
	}
	return total
}

func (p producto) CalcularCosto() float64 {
	var tipoCalculo = map[string]float64{
		"pequeño": p.precio,
		"mediano": p.precio * 1.03,
		"grande":  p.precio*1.06 + 2500,
	}
	return tipoCalculo[p.tipo]
}

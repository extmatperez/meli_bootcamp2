package main

import "fmt"

type tienda struct {
	ListaProductos []Producto
}

func (t *tienda) Total() (total float64) {
	total = 0.0
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

type tipoGrande struct{}
type tipoMediano struct{}
type tipoPequenio struct{}

func (t tipoPequenio) CalcularAdicional() func(precio float64) float64 {
	return func(precio float64) float64 {
		return precio
	}
}

func (t tipoMediano) CalcularAdicional() func(precio float64) float64 {
	return func(precio float64) float64 {
		return precio * 1.03
	}
}

func (t tipoGrande) CalcularAdicional() func(precio float64) float64 {
	funcion := func(precio float64) float64 {
		return precio*1.06 + 2500.0
	}
	return funcion
}

type producto struct {
	Tipo   tipo
	Nombre string
	Precio float64
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

func NuevoProducto(tipoProducto string, nombre string, precio float64) Producto {
	//var tip tipo
	switch tipoProducto {
	case "Grande":
		tip := tipoGrande{}
		return producto{tip, nombre, precio}
	case "Mediano":
		tip := tipoMediano{}
		return producto{tip, nombre, precio}
	case "Pequenio":
		tip := tipoPequenio{}
		return producto{tip, nombre, precio}
	}
	return producto{}
}

func NuevaTienda() Ecommerce {
	return &tienda{}
}

func main() {
	ecom := NuevaTienda()
	ecom.Agregar(NuevoProducto("Pequenio", "prodpeq", 100.0))
	fmt.Printf("Tienda 1: Total: %.2f\n", ecom.Total())
	ecom2 := NuevaTienda()
	ecom2.Agregar(NuevoProducto("Mediano", "prodmed", 100.0))
	fmt.Printf("Tienda 2: Total: %.2f\n", ecom2.Total())
	ecom3 := NuevaTienda()
	ecom3.Agregar(NuevoProducto("Grande", "prodgra", 100.0))
	fmt.Printf("Tienda 3: Total: %.2f\n", ecom3.Total())

	ecom4 := NuevaTienda()
	ecom4.Agregar(NuevoProducto("Grande", "prodgra", 100.0))
	ecom4.Agregar(NuevoProducto("Mediano", "prodgra", 100.0))
	ecom4.Agregar(NuevoProducto("Pequenio", "prodgra", 100.0))
	ecom4.Agregar(NuevoProducto("Grande", "prodgra2", 100.0))
	ecom4.Agregar(NuevoProducto("Mediano", "prodgra2", 100.0))
	ecom4.Agregar(NuevoProducto("Pequenio", "prodgra2", 100.0))
	fmt.Printf("Tienda 4: Total: %.2f\n", ecom4.Total())

}

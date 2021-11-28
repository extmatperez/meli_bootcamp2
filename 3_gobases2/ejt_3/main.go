package main

import "fmt"

func main() {
	tipoGrande := Grande{}
	tipoMediano := Mediano{}
	tipoChico := Chico{}
	prod1 := producto{tipoGrande, "Heladera", 16956.99}
	prod2 := producto{tipoMediano, "Computadora", 82365}
	prod3 := producto{tipoChico, "Celular", 67692}

	productosTienda := []producto{prod1, prod2, prod3}
	tienda1 := tienda{productosTienda}
	fmt.Printf("El costo total de los productos es de: $ %.2f\n", tienda1.Total())

	prod4 := producto{tipoGrande, "Lavarropas", 9785.99}

	tienda1.Agregar(prod4)
	fmt.Printf("El costo total de los productos es de: $ %.2f\n", tienda1.Total())
}

type tienda struct {
	Productos []producto
}

func (t *tienda) Agregar(nuevoProducto producto) {
	t.Productos = append(t.Productos, nuevoProducto)
}

func (t *tienda) Total() float64 {
	var total float64 = 0.0
	for _, prod := range t.Productos {
		fmt.Printf("El precio del producto %v es de: $%.2f\n", prod.Nombre, prod.CalcularCosto())
		total += prod.CalcularCosto()
	}
	return total
}

type producto struct {
	Tipo   Tipo
	Nombre string
	Precio float64
}

func (p *producto) CalcularCosto() float64 {
	return p.Tipo.CalcularAdicional()(p.Precio)
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(producto Producto)
}

type Tipo interface {
	CalcularAdicional() func(precio float64) float64
}

type Grande struct{}

func (n Grande) CalcularAdicional() func(precio float64) float64 {
	fun := func(precio float64) float64 {
		return precio + ((precio * 6) / 100) + 2500
	}
	return fun
}

type Mediano struct{}

func (n Mediano) CalcularAdicional() func(precio float64) float64 {
	fun := func(precio float64) float64 {
		return precio + ((precio * 3) / 100)
	}
	return fun
}

type Chico struct{}

func (n Chico) CalcularAdicional() func(precio float64) float64 {
	fun := func(precio float64) float64 {
		return precio
	}
	return fun
}

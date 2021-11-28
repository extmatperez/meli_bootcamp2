package main

import "fmt"

func main() {
	tipoGrande := Grande{}
	tipoMediano := Mediano{}
	tipoChico := Chico{}
	producto1 := producto{tipoChico, "llavero", 37.5}
	producto11 := producto{tipoChico, "llavero", 37.5}
	producto2 := producto{tipoMediano, "caja", 400}
	producto21 := producto{tipoMediano, "caja", 400}
	producto3 := producto{tipoGrande, "televisor", 8538.4}
	producto31 := producto{tipoGrande, "televisor", 8538.4}
	productosTienda := []producto{producto1, producto11, producto2, producto21, producto31, producto2, producto3}
	tienda1 := nuevaTienda(productosTienda, "tienda1")

	fmt.Printf("El costo total de los productos es de: $ %.2f\n", tienda1.total())
	//fmt.Println(tienda1)
}

type tienda struct {
	productos []producto
}

type producto struct {
	Tipo   Tipo
	nombre string
	precio float64
}

type Producto interface {
	calcularCosto() float64
}

type Ecommerce interface {
	total() float64
	agregar(producto producto)
}

func (t *tienda) total() float64 {
	total := 0.0
	for _, prod := range t.productos {
		fmt.Printf("El precio del producto %v es de: $%.2f\n", prod.nombre, prod.CalcularCosto())
		total += prod.precio
	}
	return total
}

func (t *tienda) agregar(nuevoProduct producto) {
	t.productos = append(t.productos, nuevoProduct)
}

func nuevaTienda(productosTienda []producto, tipo string) Ecommerce {
	switch tipo {
	case "tienda1":
		return &tienda{productosTienda}
	}
	return nil
}

func (p *producto) CalcularCosto() float64 {
	return p.Tipo.CalcularAdicional()(p.precio)
}

type Tipo interface {
	CalcularAdicional() func(precio float64) float64
}

// --------------------------------------------- Tipos de producto -----------------------------------------------------

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

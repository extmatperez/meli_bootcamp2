package main

import (
	"fmt"
	"math"
)

type Producto struct {
	name string
	price float64
	quantity int
}

type Servicio struct {
	name string
	price float64
	minutes int
}

type Mantenimiento struct {
	name string
	price float64
}

func sumarProductos(productos []Producto, suma chan float64) {
	var total float64
	for _, producto := range productos {
		total += producto.price * float64(producto.quantity)
	}
	suma <- total
}

func sumarServicios(servicios []Servicio, suma chan float64) {
	var total float64
	for _, servicio := range servicios {
	    mediasHoras := math.Ceil((float64(servicio.minutes)) / 30)
		total += servicio.price * mediasHoras
	}
	suma <- total
}

func sumarMantenimientos(mantenimientos []Mantenimiento, suma chan float64) {
	var total float64
	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.price
	}
	suma <- total
}

func main() {
	productos := []Producto{
		{name: "Monitor", price: 100, quantity: 2},
		{name: "Teclado", price: 50, quantity: 1},
		{name: "Mouse", price: 30, quantity: 1},
	}
	servicios := []Servicio{
		{name: "Instalación", price: 500, minutes: 60},
		{name: "Mantenimiento", price: 300, minutes: 120},
	}
	mantenimientos := []Mantenimiento{
		{name: "Instalación", price: 500},
		{name: "Mantenimiento", price: 300},
	}

	sumaProductos := make(chan float64)
	sumaServicios := make(chan float64)
	sumaMantenimientos := make(chan float64)

	go sumarProductos(productos, sumaProductos)
	go sumarServicios(servicios, sumaServicios)
	go sumarMantenimientos(mantenimientos, sumaMantenimientos)

	total := <-sumaProductos + <-sumaServicios + <-sumaMantenimientos
	fmt.Println(total)
}

/*
Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos,
Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que
el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

Se requieren 3 estructuras:
Productos: nombre, precio, cantidad.
Servicios: nombre, precio, minutos trabajados.
Mantenimiento: nombre, precio.

Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando
el total de los 3).
*/

package main

import "fmt"

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Servicio struct {
	nombre            string
	precio            float64
	minutosTrabajados int
}

type Mantenimiento struct {
	nombre string
	precio float64
}

func sumarProductos(c chan float64, productos ...Producto) {

	precioTotal := 0.0

	for _, p := range productos {
		precioTotal += p.precio * float64(p.cantidad)
	}

	c <- precioTotal
}

func sumarServicios(c chan float64, servicios ...Servicio) {

	precioTotal := 0.0

	for _, s := range servicios {

		if s.minutosTrabajados < 30 {
			precioTotal += s.precio
		} else {
			precioTotal += s.precio * (float64(s.minutosTrabajados) / 30.0)
		}
	}

	c <- precioTotal
}

func sumarMantenimientos(c chan float64, mantenimientos ...Mantenimiento) {

	precioTotal := 0.0

	for _, m := range mantenimientos {
		precioTotal += m.precio
	}

	c <- precioTotal
}

func main() {

	canalProducto := make(chan float64)
	canalServicio := make(chan float64)
	canalMantenimiento := make(chan float64)

	go sumarProductos(canalProducto, Producto{"Inflar gomas", 100, 2}, Producto{"Pinchadura", 50, 10}, Producto{"Chapa y pintura", 500, 3})
	fmt.Printf("Precio total productos: %.2f\n", <-canalProducto)

	go sumarServicios(canalServicio, Servicio{"Control de luz", 50, 65}, Servicio{"Arrego lavarropas", 80, 130})
	fmt.Printf("Precio total servicios: %.2f\n", <-canalServicio)

	go sumarMantenimientos(canalMantenimiento, Mantenimiento{"Lavada interior", 55}, Mantenimiento{"Cambio resortes", 60})
	fmt.Printf("Precio total mantenimientos: %.2f\n", <-canalMantenimiento)
}

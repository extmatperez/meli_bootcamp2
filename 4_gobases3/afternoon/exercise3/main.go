package main

import "fmt"

type Productos struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}
type Servicios struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Minutos  float64 `json:"minutos"`
	Trabajos string  `json:"trabajos"`
}
type Mantenimiento struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

func sumarProductos(p []Productos, c chan float64) {
	var suma float64
	for _, v := range p {
		suma += v.Precio * float64(v.Cantidad)
	}
	c <- suma
}
func sumarServicios(s []Servicios, c chan float64) {
	var suma float64
	for _, v := range s {
		if v.Minutos < 30 {
			suma += v.Precio * 30
		} else {
			suma += v.Precio * v.Minutos
		}
	}
	c <- suma
}
func sumarMantenimiento(m []Mantenimiento, c chan float64) {
	var suma float64
	for _, v := range m {
		suma += v.Precio
	}
	c <- suma
}
func main() {

	c := make(chan float64)
	total := 0.0

	arrayProductos := []Productos{{"pala", 10, 2}, {"hacha", 3, 10}, {"rastrillo", 4, 10}}
	arrayServicios := []Servicios{{"limpieza", 2, 30, "habitaciones"}, {"mantenimiento", 1, 30, "revision valvulas"}, {"mantenimiento", 2, 30, "cocheras"}}
	arrayMantenimiento := []Mantenimiento{{"mantenimiento", 20}, {"mantenimiento", 30}, {"mantenimiento", 20}}

	go sumarProductos(arrayProductos, c)
	go sumarServicios(arrayServicios, c)
	go sumarMantenimiento(arrayMantenimiento, c)

	for i := 0; i < 3; i++ {
		total += <-c
	}
	fmt.Println(total)

}

/*
	Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
	Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

	Se requieren 3 estructuras:
	Productos: nombre, precio, cantidad.
	Servicios: nombre, precio, minutos trabajados.
	Mantenimiento: nombre, precio.

	Se requieren 3 funciones:
	Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
	Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada, si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
	Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

	Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/

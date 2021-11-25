package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicios struct {
	Nombre        string
	Precio        float64
	MinTrabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

/*
Se requieren 3 funciones:
Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
	si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).

*/

func sumarProductos(productos []Producto, c chan float64) {
	precioTotal := 0.0
	for _, producto := range productos {
		precioTotal += producto.Precio * float64(producto.Cantidad)
	}
	c <- precioTotal
}

func sumarServicios(servicios []Servicios, c chan float64) {
	precioTotal := 0.0
	for _, servicio := range servicios {
		if servicio.MinTrabajados < 30 {
			precioTotal += servicio.Precio
		} else {
			precioTotal += servicio.Precio * float64(servicio.MinTrabajados) / 30
		}
	}
	c <- precioTotal
}

func sumarMantenimiento(mantenimientos []Mantenimiento, c chan float64) {
	precioTotal := 0.0
	for _, mantenimiento := range mantenimientos {
		precioTotal += mantenimiento.Precio
	}
	c <- precioTotal
}

func main() {
	c := make(chan float64)
	productos := []Producto{{"Alfajor", 35, 10},
		{"Chicle", 20, 50},
		{"Gomita", 55, 10},
	}

	fmt.Println(productos)

	servicios := []Servicios{{"Limpieza", 300, 60},
		{"Masajes", 700, 120},
	}

	fmt.Println(servicios)

	mantenimientos := []Mantenimiento{{"Reparacion PC", 700},
		{"Admin redes", 1300},
		{"Reparacion luz", 5000},
	}
	fmt.Println(mantenimientos)

	go sumarProductos(productos, c)
	go sumarServicios(servicios, c)
	go sumarMantenimiento(mantenimientos, c)
	valor := 0.0
	for i := 0; i < 3; i++ {
		valor += <-c
	}

	fmt.Printf("El valor total es de %.2f\n", valor)
}

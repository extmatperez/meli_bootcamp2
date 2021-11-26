package main

import "fmt"

func main() {
	producto1 := Producto{"Producto1", 56.00, 6.0}
	producto2 := Producto{"Producto2", 176.00, 2.0}
	producto3 := Producto{"Producto3", 96.00, 7.0}
	producto4 := Producto{"Producto4", 6.00, 3.0}
	producto5 := Producto{"Producto5", 489.50, 16.0}

	servicio1 := Servicio{"Servicio1", 56.00, 60}
	servicio2 := Servicio{"Servicio2", 176.00, 20}
	servicio3 := Servicio{"Servicio3", 96.00, 70}
	servicio4 := Servicio{"Servicio4", 6.00, 30}
	servicio5 := Servicio{"Servicio5", 489.50, 160}

	mantenimiento1 := Mantenimiento{"Mantenimiento1", 56.00}
	mantenimiento2 := Mantenimiento{"Mantenimiento2", 176.00}
	mantenimiento3 := Mantenimiento{"Mantenimiento3", 96.00}
	mantenimiento4 := Mantenimiento{"Mantenimiento4", 6.00}
	mantenimiento5 := Mantenimiento{"Mantenimiento5", 489.50}

	productos := []Producto{producto1, producto2, producto3, producto4, producto5}
	servicios := []Servicio{servicio1, servicio2, servicio3, servicio4, servicio5}
	mantenimientos := []Mantenimiento{mantenimiento1, mantenimiento2, mantenimiento3, mantenimiento4, mantenimiento5}

	cProductos := make(chan float64)
	go SumarProductos(productos, cProductos)
	fmt.Println("Empezo el calculo de productos")

	cServicios := make(chan float64)
	go SumarServicios(servicios, cServicios)
	fmt.Println("Empezo el calculo de servicios")

	cMantenimientos := make(chan float64)
	go SumarMantenimientos(mantenimientos, cMantenimientos)
	fmt.Println("Empezo el calculo de mantenimientos")

	fmt.Printf("Precio productos: %v\n", <-cProductos)
	fmt.Printf("Precio servicios: %v\n", <-cServicios)
	fmt.Printf("Precio mantenimientos: %v\n", <-cMantenimientos)
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad float64
}

func (p *Producto) CalcularTotal() float64 {
	return p.Precio * p.Cantidad
}

type Servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados uint32
}

func (s *Servicio) CalcularTotal() float64 {
	costo := 0.0
	if s.MinutosTrabajados >= 30 {
		mediasHoras := s.MinutosTrabajados / 30
		resto := s.MinutosTrabajados % 30
		costo += float64(mediasHoras) * s.Precio
		if resto > 0 {
			costo += s.Precio
		}
	} else if s.MinutosTrabajados > 0 {
		costo = s.Precio
	}
	return costo
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func (m *Mantenimiento) CalcularTotal() float64 {
	return m.Precio
}

type CarteraEmpresa interface {
	CalcularTotal(c chan float64)
}

func SumarProductos(productos []Producto, c chan float64) float64 {
	total := 0.0
	for _, producto := range productos {
		total += producto.CalcularTotal()
	}
	c <- total
	return <-c
}

func SumarServicios(servicios []Servicio, c chan float64) float64 {
	total := 0.0
	for _, servicio := range servicios {
		total += servicio.CalcularTotal()
	}
	c <- total
	return <-c
}

func SumarMantenimientos(mantenimientos []Mantenimiento, c chan float64) float64 {
	total := 0.0
	for _, mantenimiento := range mantenimientos {
		total += mantenimiento.CalcularTotal()
	}
	c <- total
	return <-c
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

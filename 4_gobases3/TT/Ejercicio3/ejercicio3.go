package main

import "fmt"

func main() {
	productos := []Producto{}
	servicios := []Servicio{}
	mantenimientos := []Mantenimiento{}

	p1 := Producto{"Heladera", 250, 1}
	p2 := Producto{"TV", 550, 2}
	p3 := Producto{"Lavarropas", 200, 1}
	p4 := Producto{"Sofa", 400, 2}

	productos = append(productos, p1, p2, p3, p4)

	s1 := Servicio{"Limpieza", 8, 80}
	s2 := Servicio{"Cocinar", 5, 50}
	s3 := Servicio{"Pintar", 10, 240}

	servicios = append(servicios, s1, s2, s3)

	m1 := Mantenimiento{"Piscina", 40}

	mantenimientos = append(mantenimientos, m1)

	c := make(chan float64)
	valor_final := 0.00
	go sumar_productos(productos, c)
	valor_final += <-c
	go sumar_servicios(servicios, c)
	valor_final += <-c
	go sumar_mantenimiento(mantenimientos, c)
	valor_final += <-c

	fmt.Printf("%.2f\n", valor_final)
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre             string
	Precio             float64
	Minutos_trabajados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumar_productos(productos []Producto, c chan float64) {
	suma := 0.00
	for _, p := range productos {
		suma += p.Precio * float64(p.Cantidad)
	}
	c <- suma
}

func sumar_servicios(servicios []Servicio, c chan float64) {
	suma := 0.00
	for _, s := range servicios {
		suma += s.Precio * float64((s.Minutos_trabajados/30)+1)
	}
	c <- suma
}

func sumar_mantenimiento(mantenimientos []Mantenimiento, c chan float64) {
	suma := 0.00
	for _, m := range mantenimientos {
		suma += m.Precio
	}
	c <- suma
}

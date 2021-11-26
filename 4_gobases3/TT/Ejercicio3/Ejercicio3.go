package main

import (
	"fmt"
)

type Producto struct {
	Nombre   string  `json:"nombre_producto"`
	Precio   float64 `json:"precio_producto"`
	Cantidad int     `json:"cantidad_producto"`
}

type Servicio struct {
	Nombre      string  `json:"nombre_servicio"`
	Precio      float64 `json:"precio_servicio"`
	Min_Trabajo int     `json:"min_trabajo_servicio"`
}

type Mantenimiento struct {
	Nombre string  `json:"nombre_mantenim"`
	Precio float64 `json:"precio_mantenim"`
}

func sumar_productos(c chan float64, productos []Producto) {
	var total_producto float64 = 0.0
	for _, element := range productos {
		total_producto += (element.Precio * (float64)(element.Cantidad))
	}
	c <- total_producto
}

func sumar_servicios(c chan float64, servicios []Servicio) {
	var total_servicio float64 = 0.0
	for _, element := range servicios {
		if element.Min_Trabajo%30 != 0 {
			element.Min_Trabajo = element.Min_Trabajo + ((element.Min_Trabajo % 30) + (30 - element.Min_Trabajo%30))
		}
		total_servicio += (float64(element.Min_Trabajo / 30)) * element.Precio
	}
	c <- total_servicio
}

func sumar_mantenimientos(c chan float64, mantenimientos []Mantenimiento) {
	var total_mantenim float64 = 0.0
	for _, element := range mantenimientos {
		total_mantenim += element.Precio
	}
	c <- total_mantenim
}

func main() {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	p0 := Producto{"Mopa", 112.30, 2}
	p1 := Producto{"Pala", 240.50, 1}
	p2 := Producto{"Taza", 400.00, 1}
	p3 := Producto{"Remera", 650.30, 3}

	s0 := Servicio{"Herreria", 756.00, 28}
	s1 := Servicio{"Electricidad", 560.00, 38}
	s2 := Servicio{"Orfebreria", 1400, 58}

	m0 := Mantenimiento{"Piscina", 5501.00}
	m1 := Mantenimiento{"Heladera y frigorifico", 8300.00}

	go sumar_productos(c1, []Producto{p0, p1, p2, p3})
	go sumar_servicios(c2, []Servicio{s0, s1, s2})
	go sumar_mantenimientos(c3, []Mantenimiento{m0, m1})

	var total float64 = <-c1 + <-c2 + <-c3

	fmt.Println("\nFinalizo el canal 1. ", c1)
	fmt.Println("\nFinalizo el canal 2. ", c2)
	fmt.Println("\nFinalizo el canal 3. ", c3)

	fmt.Printf("\nSumatoria total: %.2f", total)
}

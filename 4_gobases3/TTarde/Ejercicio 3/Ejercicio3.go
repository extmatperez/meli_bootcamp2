package main

import (
	"fmt"
	"time"
)

type Productos struct {
	Nombre string
	Precio float64
	Cantidad float64
}

type Servicios struct {
	Nombre string
	Precio float64
	Minutos_tr float64
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(p []Productos) float64 {
	var sum float64
	for i := 0; i < len(p); i++ {
		sum += p[i].Precio*p[i].Cantidad
	}
	return sum
}

func sumarServicios(s []Servicios) float64 {
	var sum, conv float64
	for i := 0; i < len(s); i++ {
		conv += s[i].Minutos_tr / 30
		sum += s[i].Precio * conv
	}
	return sum
}

func sumarMantenimiento(m []Mantenimiento) float64 {
var sum float64
for i := 0; i < len(m);i++ {
	sum += m[i].Precio
}
return sum
}




func main() {
	p1 := Productos{Nombre: "Juan", Precio: 244.65, Cantidad: 5}
	p2 := Productos{Nombre:"Pedro", Precio: 214.88, Cantidad: 9}

	s1 := Servicios{Nombre:"Luz", Precio: 100, Minutos_tr: 55}
	s2 := Servicios{Nombre:"Gas", Precio: 160, Minutos_tr: 35}

	m1 := Mantenimiento{Nombre:"Reparacion", Precio:544}
	m2 := Mantenimiento{Nombre:"Actualizacion", Precio:437}

	list := make([]Productos, 0)
	list = append(list, p1, p2)
	list2 := make([]Servicios, 0)
	list2 = append(list2, s1, s2)
	list3 := make([]Mantenimiento, 0)
	list3 = append(list3,m1,m2)

	begins := time.Now()
	go sumarProductos(list)
	go sumarServicios(list2)
	go sumarMantenimiento(list3)

	ends := time.Now()

	total := ends.Sub(begins)
	fmt.Println(total)

}
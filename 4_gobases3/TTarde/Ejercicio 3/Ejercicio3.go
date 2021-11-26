package main

import (
	"fmt"
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

func sumarProductos(p []Productos, c chan float64){
	var sum float64
	for i := 0; i < len(p); i++ {
		sum += p[i].Precio*p[i].Cantidad
	}
	fmt.Println("Listo")
	c<-sum
}

func sumarServicios(s []Servicios, c chan float64) {
	var sum, conv float64
	for i := 0; i < len(s); i++ {
		conv += s[i].Minutos_tr / 30
		sum += s[i].Precio * conv
	}
	fmt.Println("Listo")
	c<-sum
}

func sumarMantenimiento(m []Mantenimiento, c chan float64){
var sum float64
for i := 0; i < len(m);i++ {
	sum += m[i].Precio
}
	fmt.Println("Listo")
	c<-sum
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

	c := make(chan float64)
	c1 := make(chan float64)
	c2 := make(chan float64)

	


	go sumarProductos(list, c)
	variable := <- c
	go sumarServicios(list2, c1)
	variable1 := <- c1
	go sumarMantenimiento(list3, c2)
	variable2 := <- c2

	fmt.Printf("La suma de los totales es: %v", variable+variable1+variable2)

}
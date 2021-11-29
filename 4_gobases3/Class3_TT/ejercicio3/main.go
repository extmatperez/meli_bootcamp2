package main

import "fmt"

type producto struct {
	nombre   string  //
	precio   float64 //
	cantidad int
}

type servicios struct {
	nombre   string  //
	precio   float64 //
	min_trab int     //
}

type mantenimiento struct {
	nombre string  //
	precio float64 //
}

func sumar_productos(productos []producto, c chan float64) {
	total_prod := 0.0
	for _, p := range productos {
		pre := p.precio * float64(p.cantidad)
		total_prod += pre
	}
	c <- total_prod
}

func sumar_servicios(servicios []servicios, c chan float64) {
	total_servicios := 0.0
	for _, s := range servicios {
		pre := s.precio * (float64(s.min_trab) / 60)
		total_servicios += pre

	}
	c <- total_servicios
}

func sumar_mantenimiento(mantenimiento []mantenimiento, c chan float64) {
	total_mant := 0.0
	for _, m := range mantenimiento {
		total_mant += m.precio
	}
	c <- total_mant
}

func main() {
	c := make(chan float64)

	p1 := producto{"arbol", 435.6, 12}
	p2 := producto{"caña", 755.6, 4}

	productos_ := []producto{p1, p2}
	go sumar_productos(productos_, c)
	fmt.Println(productos_)
	//fmt.Println(total_prod)

	s1 := servicios{"limpieza", 850, 120}
	s2 := servicios{"pintado", 2300, 90}

	servicios_ := []servicios{s1, s2}
	go sumar_servicios(servicios_, c)
	fmt.Println(servicios_)
	//fmt.Println(total_serv)

	m1 := mantenimiento{"mecanica", 2900}
	m2 := mantenimiento{"plomeria", 4322}

	mantenimientos_ := []mantenimiento{m1, m2}
	go sumar_mantenimiento(mantenimientos_, c)
	fmt.Println(mantenimientos_)
	//fmt.Println(total_mant)

	valor := 0.0
	for i := 0; i < 3; i++ {
		valor += <-c
	}
	fmt.Printf("El valor total es de %.2f\n", valor)
}

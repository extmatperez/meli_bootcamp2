package main

import (
	"fmt"
	"time"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre  string
	Precio  float64
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(productos []Producto, c chan float64) {
	total := 0.0
	for _, prod := range productos {
		total += prod.Precio * float64(prod.Cantidad)
	}
	c <- total
}

func sumarServicios(servicios []Servicio, c chan float64) {
	total := 0.0
	for _, serv := range servicios {
		tiempo := serv.Minutos / 30.0
		if tiempo > 1.0 {
			total += serv.Precio * float64(serv.Minutos)
		} else {
			total += serv.Precio * 30.0
		}
	}
	c <- total
}

func sumarMantenimiento(mantenimiento []Mantenimiento, c chan float64) {
	total := 0.0
	for _, mant := range mantenimiento {
		total += mant.Precio
	}
	c <- (total)
}

func main() {
	p1 := Producto{"Leche", 657.0, 3}
	p2 := Producto{"Manzana", 234.0, 7}
	p3 := Producto{"Queso", 123.0, 9}
	p4 := Producto{"Pan", 63.0, 12}

	prodList := []Producto{p1, p2, p3, p4}
	//fmt.Println(sumarProductos(prodList))

	s1 := Servicio{"S1", 10.0, 10}
	s2 := Servicio{"S2", 700.0, 35}
	s3 := Servicio{"S3", 23.0, 64}
	s4 := Servicio{"S4", 45.0, 32}

	servList := []Servicio{s1, s2, s3, s4}

	//fmt.Println(sumarServicios(servList))

	m1 := Mantenimiento{"M1", 10.0}
	m2 := Mantenimiento{"M2", 700.0}
	m3 := Mantenimiento{"M3", 23.0}
	m4 := Mantenimiento{"M4", 45.0}

	mantServ := []Mantenimiento{m1, m2, m3, m4}

	c := make(chan float64)
	ini := time.Now()
	go sumarProductos(prodList, c)
	go sumarServicios(servList, c)
	go sumarMantenimiento(mantServ, c)

	total := 0.0
	for i := 0; i < 3; i++ {
		variable := <-c
		total += variable
		fmt.Println(variable)
	}
	fin := time.Now()
	tiempo := fin.Sub(ini)
	fmt.Println(tiempo)
	fmt.Println(total)
}

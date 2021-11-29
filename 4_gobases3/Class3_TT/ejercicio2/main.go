package main

import (
	"encoding/json"
	"fmt"
)

type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type usuario struct {
	nombre    string
	apellido  string
	productos []producto
}

func NuevoProducto(nombre string, precio float64) producto {
	prod := producto{nombre: nombre, precio: precio}
	return prod
}

func AgregarProducto(us *usuario, p *producto, cant int) {
	p.cantidad = cant
	us.productos = append(us.productos, *p)
}

func BorrarProducto(us *usuario) {
	us.productos = []producto{}
}

func main() {
	u1 := usuario{nombre: "jose", apellido: "perez"}

	p1 := NuevoProducto("jabon", 255.3)
	AgregarProducto(&u1, &p1, 4)
	AgregarProducto(&u1, &p1, 8)

	//p2 := NuevoProducto("desodorante", 465.3)
	fmt.Println(u1)

	BorrarProducto(&u1)
	fmt.Println(u1)

	usuarios := []usuario{u1}
	usrForm, err := json.Marshal(usuarios)

	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(string(usrForm))
	}
}

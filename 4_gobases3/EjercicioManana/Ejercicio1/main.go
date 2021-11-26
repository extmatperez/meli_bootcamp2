package main

import (
	"fmt"
	"os"
)

type Producto struct {
	Id       string
	Precio   float64
	Cantidad int
}

func guardarEnArchivo(nombre string, p []Producto) {

	contenido := "Id;Precio;Cantidad\n"

	for _, valor := range p {
		contenido = contenido + fmt.Sprintf("%s;%.2f;%d\n", valor.Id, valor.Precio, valor.Cantidad)
	}
	fileString := []byte(contenido)
	os.WriteFile("./"+nombre, fileString, 0644)

}

func main() {

	p1 := Producto{"SVN-1", 1000, 100}
	p2 := Producto{"SVN-2", 6000, 200}
	p3 := Producto{"SVN-3", 9000, 50}
	p4 := Producto{"SVN-4", 12000, 1}
	p5 := Producto{"SVN-5", 12, 10}

	var productos []Producto
	productos = append(productos, p1)
	productos = append(productos, p2)
	productos = append(productos, p3)
	productos = append(productos, p4)
	productos = append(productos, p5)

	guardarEnArchivo("prueba1.csv", productos)

}

package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

type Producto struct {
	Nombre   string
	precio   float64
	cantidad int
}

func nuevoProducto(nombreProducto string, precioProducto float64) Producto {
	nuevoProducto := Producto{nombreProducto, precioProducto, 5}
	return nuevoProducto
}

func agregarProducto(usuario *Usuario, producto Producto) {
	usuario.Productos = append(usuario.Productos, producto)
}

func eliminarProductos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}

func main() {
	usuario1 := Usuario{
		Nombre:   "Damian",
		Apellido: "Zamora",
		Correo:   "damian.zamora@hotmail.com",
	}
	producto1 := Producto{"Esponja", 40.50, 6}
	producto2 := nuevoProducto("Escoba", 35.00)
	agregarProducto(&usuario1, producto1)
	agregarProducto(&usuario1, producto2)
	fmt.Println(usuario1)
	fmt.Println(producto2)
	eliminarProductos(&usuario1)
	fmt.Println(usuario1)
}

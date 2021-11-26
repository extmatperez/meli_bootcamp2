package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Usuario struct {
	Nombre, Apellido, Correo string
	Productos                []Producto
}

func nuevoProducto(nombre string, precio float64) (p *Producto) {
	prod := Producto{Nombre: nombre, Precio: precio}
	p = &prod
	return
}

func agregarUsuario(usuario *Usuario, producto *Producto, cantidad int) {
	(*producto).Cantidad = cantidad
	(*usuario).Productos = append((*usuario).Productos, *producto)
}

func borrarProdcto(u *Usuario) {
	(*u).Productos = []Producto{}
}

func main() {
	user := Usuario{"Nicolas", "Aponte", "nico@correo.com", []Producto{}}
	userP := &user

	prod := nuevoProducto("Chocolate", 45.0)
	agregarUsuario(userP, prod, 5)

	prod = nuevoProducto("Leche", 55.0)
	agregarUsuario(userP, prod, 5)

	prod = nuevoProducto("Café", 30.0)
	agregarUsuario(userP, prod, 10)

	fmt.Println(user)

	borrarProdcto(userP)

	fmt.Println(user)

}

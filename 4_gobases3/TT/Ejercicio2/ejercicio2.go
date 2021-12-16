package main

import "fmt"

func main() {
	usuario := Usuario{Nombre: "Pedro", Apellido: "Garcia", Correo: "pg@gmail.com"}
	prod := nuevo_producto("Lavarropas", 350)
	agregar_producto(&usuario, &prod, 2)
	fmt.Printf("%v\n", usuario.Productos)
	borrar_productos(&usuario)
	fmt.Printf("%v\n", usuario.Productos)
}

type Usuario struct {
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Correo    string     `json:"correo"`
	Productos []Producto `json:"producto"`
}

type Producto struct {
	Nombre   string `json:"nombre"`
	Precio   int    `json:"precio"`
	Cantidad int    `json:"cantidad"`
}

func nuevo_producto(nombre string, precio int) Producto {
	return Producto{Nombre: nombre, Precio: precio}
}

func agregar_producto(usuario *Usuario, producto *Producto, cantidad int) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}

func borrar_productos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}

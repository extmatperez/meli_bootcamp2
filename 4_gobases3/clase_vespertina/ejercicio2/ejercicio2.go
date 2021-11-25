package main

import "fmt"

func main() {
	usuario1 := Usuario{"Andres", "Ghione", "ag@gmail.com", []Producto{}}
	fmt.Printf("%v\n", usuario1)

	producto := NuevoProducto("Heladera", 15678)
	AgregarProducto(&usuario1, &producto, 7)
	fmt.Printf("%v\n", usuario1)
	producto = NuevoProducto("Mouse", 1068)
	AgregarProducto(&usuario1, &producto, 1)
	producto = NuevoProducto("Televisor", 29628)
	AgregarProducto(&usuario1, &producto, 1)
	producto = NuevoProducto("Comoda", 9078)
	AgregarProducto(&usuario1, &producto, 2)
	producto = NuevoProducto("Ropero", 21300)
	AgregarProducto(&usuario1, &producto, 1)
	producto = NuevoProducto("Teclado", 1800)
	AgregarProducto(&usuario1, &producto, 1)
	fmt.Printf("%v\n", usuario1)

	BorrarProductos(&usuario1)
	fmt.Printf("%v\n", usuario1)
}

type Usuario struct {
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Correo    string     `json:"correo"`
	Productos []Producto `json:"productos"`
}

type Producto struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad float64 `json:"cantidad"`
}

func NuevoProducto(nombre string, precio float64) Producto {
	producto := Producto{Nombre: nombre, Precio: precio}
	return producto
}

func AgregarProducto(usuario *Usuario, producto *Producto, cantidad float64) {
	producto.Cantidad = cantidad
	usuario.Productos = append(usuario.Productos, *producto)
}

func BorrarProductos(usuario *Usuario) {
	usuario.Productos = []Producto{}
}

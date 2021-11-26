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
	Precio   float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) Producto {
	P := Producto{Nombre: nombre, Precio: precio}
	return P
}

func agregarProducto(u *Usuario, p Producto, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, p)
}

func borrarProductos(u *Usuario) {
	u.Productos = nil
}

func main() {

	var ProductosEmpresa []Producto
	ProductosEmpresa = append(ProductosEmpresa, nuevoProducto("Producto1", 100.0))
	ProductosEmpresa = append(ProductosEmpresa, nuevoProducto("Producto2", 200.0))
	ProductosEmpresa = append(ProductosEmpresa, nuevoProducto("Producto3", 300.0))
	ProductosEmpresa = append(ProductosEmpresa, nuevoProducto("Producto4", 400.0))

	fmt.Printf("\nProductos: %+v\n", ProductosEmpresa)

	ProductosEmpresa = append(ProductosEmpresa, nuevoProducto("Producto5", 500.0))

	fmt.Printf("\nAgrendando nuevo Producto: %+v\n", ProductosEmpresa)

	usuarios := []Usuario{{Nombre: "user1", Apellido: "apellido1", Correo: "correo1"}, {Nombre: "user2", Apellido: "apellido2", Correo: "correo2"}, {Nombre: "user3", Apellido: "apellido3", Correo: "correo3"}}

	fmt.Printf("\nUsuarios: %+v\n", usuarios)

	agregarProducto(&usuarios[0], ProductosEmpresa[1], 10)
	agregarProducto(&usuarios[0], ProductosEmpresa[0], 20)
	agregarProducto(&usuarios[0], ProductosEmpresa[2], 1)

	fmt.Printf("\nusuario0: %+v\n", usuarios[0])

	fmt.Printf("\nTodoslos usuarios: %+v\n", usuarios)

	borrarProductos(&usuarios[0])

	fmt.Printf("\nborrando productos usuario0: %+v\n", usuarios[0])

}

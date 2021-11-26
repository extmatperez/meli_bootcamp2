package main

import "fmt"

type producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []producto
}

func nuevoProducto(nombre string, precio float64) producto {
	productoNuevo := producto{nombre, precio, 1}
	return productoNuevo
}

func (u *usuario) agregarProducto(p producto) {
	u.Productos = append(u.Productos, p)
}

func borrarProductos(u *usuario) {
	(*u).Productos = nil
}

func main() {

	fmt.Println("Bienvenido al ejercicio 2")

	usuario1 := usuario{"Pato", "Pallua", "mail1@mail.com", nil}
	fmt.Println("Se creo al usuario1", usuario1)

	producto1 := nuevoProducto("Producto1", 1500.0)
	fmt.Println("Se creo el producto1", producto1)

	producto2 := nuevoProducto("Producto2", 2500.0)
	fmt.Println("Se creo el producto2", producto2)

	producto3 := nuevoProducto("Producto3", 3500.0)
	fmt.Println("Se creo el producto3", producto3)

	fmt.Println("Agregamos un producto al usuario1")
	usuario1.agregarProducto(producto2)
	fmt.Println(usuario1)

	fmt.Println("Agregamos un producto al usuario1")
	usuario1.agregarProducto(producto1)
	fmt.Println(usuario1)

	fmt.Println("Quitamos un producto al usuario1")
	borrarProductos(&usuario1)
	fmt.Println(usuario1)
}

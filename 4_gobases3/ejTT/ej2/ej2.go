package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}
type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []Producto
}

func nuevoProducto(nombre string, precio float64) Producto {
	return Producto{Nombre: nombre, Precio: precio}
}

func agregarProducto(user *Usuario, producto Producto, cantidad int) {
	producto.Cantidad = cantidad
	user.Productos = append(user.Productos, producto)
}
func borrarProductos(user *Usuario) {
	user.Productos = make([]Producto, 0)
}

func main() {
	cuchara := nuevoProducto("cuchara", 30.4)
	tenedor := nuevoProducto("tenedor", 22.0)
	p := Usuario{
		Nombre: "Juan",
	}
	agregarProducto(&p, cuchara, 2)
	agregarProducto(&p, cuchara, 7)
	fmt.Println("productos:", p.Productos)
	borrarProductos(&p)
	fmt.Println("productos:", p.Productos)
	agregarProducto(&p, tenedor, 7)
	fmt.Println("productos:", p.Productos)
}

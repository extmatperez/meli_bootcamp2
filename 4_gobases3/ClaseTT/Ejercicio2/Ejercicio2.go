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
	return Producto{Nombre: nombre, Precio: precio}
}
func agregarProducto(user *Usuario, producto *Producto, cantidad int) {
	(*producto).Cantidad = cantidad
	(*user).Productos = append((*user).Productos, *producto)
}
func borrarProducto(user *Usuario) {
	(*user).Productos = []Producto{}
}
func main() {
	user := Usuario{Nombre: "Nahuel", Apellido: "SC", Correo: "ns@gmail.com"}
	p1 := nuevoProducto("Cepillo", 25.50)
	fmt.Printf("%+v \n", p1)

	agregarProducto(&user, &p1, 2)
	p2 := nuevoProducto("Cargador", 22.50)

	agregarProducto(&user, &p2, 2)

	fmt.Printf("%+v \n", user)

	borrarProducto(&user)

	fmt.Printf("%+v \n", user)

}

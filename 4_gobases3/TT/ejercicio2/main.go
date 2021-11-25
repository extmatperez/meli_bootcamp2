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
	Productos []*Producto
}

func main() {
	user := Usuario{"Seba", "Chiappa", "Correo", nil}
	MostrarUsuario(&user)
	producto := NuevoProducto("Computadora", 12.3)
	AgregarProducto(&user, producto, 3)
	producto = NuevoProducto("Celular", 133.3)
	AgregarProducto(&user, producto, 5)
	MostrarUsuario(&user)
	BorrarProductos(&user)
	MostrarUsuario(&user)
}

func NuevoProducto(nombre string, precio float64) *Producto {
	producto := Producto{nombre, precio, 0}
	return &producto
}

func AgregarProducto(user *Usuario, prod *Producto, cantidad int) {
	prod.Cantidad = cantidad
	user.Productos = append(user.Productos, prod)
}
func BorrarProductos(user *Usuario) {
	user.Productos = user.Productos[:0]
}

func MostrarUsuario(usr *Usuario) {
	fmt.Printf("Nombre: %s %s\nCorreo: %s\n", usr.Nombre, usr.Apellido, usr.Correo)
	productos := usr.Productos
	for i, producto := range productos {
		fmt.Printf("Producto %d:\n", i+1)
		fmt.Printf("> Nombre: %s\n", producto.Nombre)
		fmt.Printf("> Precio: %0.2f\n", producto.Precio)
		fmt.Printf("> Cantidad: %d\n\n", producto.Cantidad)
	}
}

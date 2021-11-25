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
	cantidad int
}

func NuevoProducto(nombre string, precio float64) Producto {
	p1 := Producto{
		Nombre:   nombre,
		Precio:   precio,
		cantidad: 1,
	}
	return p1
}

func AgregarProducto(us *Usuario, pr Producto, cant int) {
	(*us).Productos = append((*us).Productos, pr)

}
func BorrarProductos(us *Usuario) {
	(*us).Productos = nil

}

func main() {
	pr1 := NuevoProducto("mouse", 70000.00)
	fmt.Printf("Producto 1: %+v\n", pr1)
	pr2 := NuevoProducto("teclado", 155600.10)
	fmt.Printf("Producto 2: %+v\n", pr2)

	us1 := Usuario{
		Nombre:   "Diego",
		Apellido: "Parra",
		Correo:   "dparra@correo.com",
	}
	AgregarProducto(&us1, pr1, 5)
	AgregarProducto(&us1, pr2, 2)

	fmt.Printf("Usuario 1: %+v\n", us1)

	BorrarProductos(&us1)

	fmt.Printf("Usuario 1: %+v\n", us1)

}

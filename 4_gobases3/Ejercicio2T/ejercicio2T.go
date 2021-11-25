package main

import "fmt"

type Usuario struct {
	Nombre, Apellido, Correo string
	ListProductos            []Producto
}

func (u *Usuario) setNombre(nombre string) {
	u.Nombre = nombre
}
func (u *Usuario) setApellido(apellido string) {
	u.Apellido = apellido
}
func (u *Usuario) setCorreo(correo string) {
	u.Correo = correo
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func nuevoProducto(nombre string, precio float64) Producto {
	p := Producto{
		Nombre: nombre,
		Precio: precio,
	}
	return p
}

func agregarProducto(c int, u *Usuario, p Producto) {
	p.Cantidad = c
	(*u).ListProductos = append((*u).ListProductos, p)

}
func (u *Usuario) borrarProducto() {
	u.ListProductos = nil
}
func main() {
	fmt.Println()

	var usuario1 Usuario
	usuario1.setNombre("Walter")
	usuario1.setApellido("Castillo")
	usuario1.setCorreo("walter@walter.com")
	fmt.Println(usuario1)

	producto1 := nuevoProducto("arroz", 10.00)
	producto2 := nuevoProducto("azucar", 20.00)
	fmt.Println(producto1)

	fmt.Println()

	agregarProducto(2, &usuario1, producto1)
	agregarProducto(2, &usuario1, producto2)
	fmt.Println(usuario1)

	fmt.Println()

	usuario1.borrarProducto()
	fmt.Println(usuario1)

	fmt.Println()

}

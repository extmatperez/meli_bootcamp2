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

func nuevoProducto(nom string, precio float64) Producto {
	newProd := Producto{nom, precio, 1}

	return newProd
}

func (u *Usuario) agregarProducto(prod Producto, cant int) {
	newProd := Producto{prod.Nombre, prod.Precio, cant}
	u.Productos = append(u.Productos, newProd)
}

func (u *Usuario) borrarProducto() {
	u.Productos = make([]Producto, 0)
}

func main() {
	var listaProd []Producto
	prod := nuevoProducto("jabon", 30.00)
	prod2 := nuevoProducto("shampoo", 60.00)
	usuario := Usuario{"ivan", "arevalo", "gol", listaProd}
	usuario.agregarProducto(prod, 3)
	usuario.agregarProducto(prod2, 5)
	fmt.Printf("%v\n", usuario)

}

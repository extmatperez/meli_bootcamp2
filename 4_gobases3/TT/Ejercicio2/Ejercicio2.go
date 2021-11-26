package main

import "fmt"

type Producto struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

type Usuario struct {
	Nombre    string     `json:"nombre"`
	Apellido  string     `json:"apellido"`
	Correo    string     `json:"correo"`
	Productos []Producto `json:"productos"`
}

func (p *Producto) nuevo_producto(nombre string, precio float64) Producto {
	return Producto{nombre, precio, 1}
}

func (u *Usuario) agregar_producto(producto *Producto, cantidad int) Usuario {
	u.Productos = append(u.Productos, Producto{producto.Nombre, producto.Precio, cantidad})
	return Usuario{u.Nombre, u.Apellido, u.Correo, u.Productos}
}

func (u *Usuario) borrar_productos() Usuario {
	var producto_vacio []Producto
	u.Productos = producto_vacio
	return Usuario{u.Nombre, u.Apellido, u.Correo, u.Productos}
}

func main() {
	// Que nos muestre un nuevo producto.
	var p0 Producto
	p0 = p0.nuevo_producto("Cereales", 120.45)
	p1 := Producto{"Arvejas", 30.50, 4}
	p2 := Producto{"Choclo", 5.60, 5}
	u0 := Usuario{"Rodrigo", "Vega Gimenez", "rvega389@gmail.com", []Producto{p1, p2}}
	fmt.Println(u0)
	u0.agregar_producto(&p0, 4)
	fmt.Println(u0)
	u0.borrar_productos()
	fmt.Println(u0)
}

package main

import "fmt"

func main() {
	var p0 Producto
	p0 = p0.nuevoProducto("Cereales", 120)
	p1 := Producto{"Arvejas", 30, 4}
	p2 := Producto{"Choclo", 5, 5}
	u0 := Usuario{"Rodrigo", "Romero", "rir@gmail.com", []Producto{p1, p2}}
	fmt.Println(u0)
	u0.agregarProducto(&p0, 4)
	fmt.Println(u0)
	u0.borrarProductos()
	fmt.Println(u0)

}

type Usuario struct {
	nombre, apellido, correo string
	productos                []Producto
}

type Producto struct {
	nombre           string
	precio, cantidad int
}

func (p *Producto) nuevoProducto(nombre string, precio int) Producto {
	prod := Producto{nombre, precio, 1}
	return prod
}

func (u *Usuario) agregarProducto(prod *Producto, cant int) Usuario {
	u.productos = append(u.productos, Producto{prod.nombre, prod.precio, cant})
	return Usuario{u.nombre, u.apellido, u.correo, u.productos}
}

func (u *Usuario) borrarProductos() Usuario {
	var listavacia []Producto
	u.productos = listavacia
	return Usuario{u.nombre, u.apellido, u.correo, u.productos}
}

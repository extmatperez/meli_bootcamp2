package main

import(
	"fmt"
)

type Producto struct{
	Nombre string
	Precio float64
	Cantidad int
}


type Usuario struct {
	Nombre string
	Apellido string
	Correo string
	Productos []Producto
}

func nuevoProducto(nombre string, precio float64) Producto{
	return Producto{nombre, precio, 0}
}

func (u *Usuario)agregarProducto(p *Producto, cantidad int){
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, *p)
}

func (u *Usuario) borrarProducto(){
	u.Productos = []Producto{}
}


func main(){
	fmt.Println("Terminado")
}
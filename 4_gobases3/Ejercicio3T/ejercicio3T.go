package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}
type Servicio struct {
	Nombre        string
	Precio        float64
	MinTrabajados int
}
type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(list []Producto, c chan float64) { //float64 {
	var aux float64
	for _, p := range list {
		aux += (p.Precio * float64(p.Cantidad))
	}
	c <- aux
	//fmt.Println("El precio total de todos los productos es:", aux)
	//return aux
}
func SumarServicios(list []Servicio, c chan float64) { // float64 {
	var aux float64
	for _, s := range list {
		if s.MinTrabajados < 30 {
			s.MinTrabajados = 30
		}
		tiempoTrabajado := float64(s.MinTrabajados / 30)
		aux += (s.Precio * tiempoTrabajado)
	}
	c <- aux
	//fmt.Println("El precio total de todos los servicios es:", aux)
	//return aux
}
func SumarMantenimientos(list []Mantenimiento, c chan float64) { // float64 {
	var aux float64
	for _, m := range list {
		aux += m.Precio
	}
	c <- aux
	//fmt.Println("El precio total de todos los mantenimientos es:", aux)
	//return aux
}
func main() {
	fmt.Println()
	var ListProductos []Producto
	var ListServicios []Servicio
	var ListMantenimientos []Mantenimiento
	ListProductos = append(ListProductos, Producto{"arroz", 10.10, 2}, Producto{"azucar", 20.00, 3}, Producto{"pan", 50, 10},
		Producto{"pan", 50, 10}, Producto{"pan", 50, 10}, Producto{"pan", 50, 10})
	//Producto{"pan", 50, 10}, Producto{"pan", 50, 10}, Producto{"pan", 50, 10})
	ListServicios = append(ListServicios, Servicio{"pintada", 10, 1}, Servicio{"lavada", 20, 60})
	ListMantenimientos = append(ListMantenimientos, Mantenimiento{"carro", 1000.00}, Mantenimiento{"moto", 500})

	c := make(chan float64)
	go SumarProductos(ListProductos, c)
	go SumarServicios(ListServicios, c)
	go SumarMantenimientos(ListMantenimientos, c)
	//fmt.Println("El precio total de todos los servicios es:", <-c)
	//fmt.Println("El precio total de todos los Mantenimientos es:", <-c)
	//fmt.Println("El precio total de todos los productos es:", <-c)
	aux := 0.0
	for i := 0; i < 3; i++ {
		v := <-c
		fmt.Println(v)
		aux += v
	}
	fmt.Println("El total es:", aux)
	fmt.Println()
}

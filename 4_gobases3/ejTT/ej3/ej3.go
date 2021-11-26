package main

import "fmt"

// CON LAS LINEAS COMENTADAS SE HARIA EL PROBLEMA
// CALCULANDO INDIVIDUALMENTE CADA SERVICIO, PRESTAMO Y MANTENIMIENTO

type Producto struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}
type Servicio struct {
	Nombre            string  `json:"nombre"`
	Precio            float64 `json:"precio"`
	MinutosTrabajados int     `json:"minutosTrabajados"`
}
type Mantenimiento struct {
	Nombre string  `json:"nombre"`
	Precio float64 `json:"precio"`
}

func sumarProductos(productos []Producto, c chan float64) {
	sumatoria := 0.0
	for _, p := range productos {
		sumatoria += p.Precio * float64(p.Cantidad)
	}
	c <- sumatoria

}
func sumarServicios(servicios []Servicio, c chan float64) {
	sumatoria := 0.0
	for _, s := range servicios {
		var minutos int
		if s.MinutosTrabajados < 30 {
			minutos = 30
		}
		minutos = s.MinutosTrabajados / 30
		sumatoria += s.Precio * float64(minutos)
	}
	c <- sumatoria

}
func sumarMantenimientos(mantenimientos []Mantenimiento, c chan float64) {
	sumatoria := 0.0
	for _, m := range mantenimientos {
		sumatoria += float64(m.Precio)
	}
	c <- sumatoria

}

// func calcularSumaTotal(valores ...float64) float64 {
// 	sum := 0.0
// 	for _, val := range valores {
// 		sum += val
// 	}
// 	return sum
// }

func main() {
	producto1 := Producto{"Producto1", 56.00, 6.0}
	producto2 := Producto{"Producto2", 176.00, 2.0}
	producto3 := Producto{"Producto3", 96.00, 7.0}
	producto4 := Producto{"Producto4", 6.00, 3.0}
	producto5 := Producto{"Producto5", 489.50, 16.0}

	servicio1 := Servicio{"Servicio1", 56.00, 60}
	servicio2 := Servicio{"Servicio2", 176.00, 20}
	servicio3 := Servicio{"Servicio3", 96.00, 70}
	servicio4 := Servicio{"Servicio4", 6.00, 30}
	servicio5 := Servicio{"Servicio5", 489.50, 160}

	mantenimiento1 := Mantenimiento{"Mantenimiento1", 56.00}
	mantenimiento2 := Mantenimiento{"Mantenimiento2", 176.00}
	mantenimiento3 := Mantenimiento{"Mantenimiento3", 96.00}
	mantenimiento4 := Mantenimiento{"Mantenimiento4", 6.00}
	mantenimiento5 := Mantenimiento{"Mantenimiento5", 489.50}

	productos := []Producto{producto1, producto2, producto3, producto4, producto5}
	servicios := []Servicio{servicio1, servicio2, servicio3, servicio4, servicio5}
	mantenimientos := []Mantenimiento{mantenimiento1, mantenimiento2, mantenimiento3, mantenimiento4, mantenimiento5}
	a := make(chan float64)
	// b := make(chan float64)
	// c := make(chan float64)
	go sumarProductos(productos, a)
	go sumarServicios(servicios, a)
	go sumarMantenimientos(mantenimientos, a)

	total_productos := <-a
	total_productos += <-a
	total_productos += <-a

	// total_servicios := <-b
	// total_mantenimientos := <-c
	fmt.Println("suma parcial productos:", total_productos)
	// fmt.Println("suma parcial servicios:", total_servicios)
	// fmt.Println("suma parcial mantenimientos:", total_mantenimientos)
	// suma_total := calcularSumaTotal(total_productos, total_servicios, total_mantenimientos)
	// fmt.Println("total:", suma_total)
}

package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}
type Servicio struct {
	Nombre            string
	Precio            float64
	MinutosTrabajados int
}
type Mantenimiento struct {
	Nombre string
	Precio float64
}

func main() {
	var productos []Producto
	var servicios []Servicio
	var mantenimeintos []Mantenimiento
	productos = append(productos, Producto{"p1", 123.333, 3}, Producto{"p2", 1233.333, 5}, Producto{"p3", 1663.666, 30})
	servicios = append(servicios, Servicio{"s1", 11.2, 62}, Servicio{"s2", 100.0, 118}, Servicio{"s3", 15.5, 7})
	mantenimeintos = append(mantenimeintos, Mantenimiento{"m1", 11.2}, Mantenimiento{"m2", 15.7})
	cp := make(chan float64)
	cs := make(chan float64)
	cm := make(chan float64)
	go SumarProductos(productos, cp)
	go SumarServicios(servicios, cs)
	go SumarMantenimientos(mantenimeintos, cm)

	// fmt.Printf("El total de productos es $%0.2f\n", <-cp)
	// fmt.Printf("El total de servicios es $%0.2f\n", <-cs)
	// fmt.Printf("El total de mantenimientos es $%0.2f\n", <-cm)
	fmt.Printf("El total a pagar es $%0.2f\n", <-cp+<-cs+<-cm)

}

func SumarProductos(productos []Producto, c chan float64) {
	total := 0.0
	for _, producto := range productos {
		total += producto.Precio * float64(producto.Cantidad)
	}
	c <- total
}

func SumarServicios(servicios []Servicio, c chan float64) {
	total := 0.0
	for _, servicio := range servicios {
		mediasHoras := servicio.MinutosTrabajados / 30
		if servicio.MinutosTrabajados%30 > 0 {
			mediasHoras++
		}
		total += servicio.Precio * float64(mediasHoras)
	}
	c <- total
}

func SumarMantenimientos(mantenimeintos []Mantenimiento, c chan float64) {
	total := 0.0
	for _, manten := range mantenimeintos {
		total += manten.Precio
	}
	c <- total
}

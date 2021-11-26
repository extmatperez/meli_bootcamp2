package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre  string
	Precio  float64
	Minutos int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func sumarProductos(p []Producto, canal chan float64) {
	precioTotal := 0.0

	for _, valor := range p {
		fmt.Printf("\n productos precio: %f * cantidad: %d\n", valor.Precio, valor.Cantidad)
		precioTotal += valor.Precio * float64(valor.Cantidad)
	}

	canal <- precioTotal
}

func sumarServicios(s []Servicio, canal chan float64) {
	precioTotal := 0.0

	for _, valor := range s {
		if valor.Minutos%30 > 0 {
			fmt.Printf("\n servicio precio: %f * hora: %f \n", valor.Precio, float64((valor.Minutos/30)+1))
			precioTotal += valor.Precio * float64((valor.Minutos/30)+1)
		} else {
			fmt.Printf("\n servicio precio: %f * hora: %f \n", valor.Precio, float64(valor.Minutos/30))
			precioTotal += valor.Precio * float64(valor.Minutos/30)
		}

	}

	canal <- precioTotal
}

func sumarMantenimiento(m []Mantenimiento, canal chan float64) {
	precioTotal := 0.0

	for _, valor := range m {
		fmt.Printf("\n mantenimiento precio: %f\n", valor.Precio)
		precioTotal += valor.Precio
	}

	canal <- precioTotal
}

func main() {

	productos := []Producto{{"Producto1", 111.0, 10}, {"Producto2", 12.0, 20}, {"Producto3", 13.0, 30}, {"Producto4", 14.0, 40}, {"Producto5", 55.0, 5}}
	servicios := []Servicio{{"S1", 100, 30}, {"S2", 200, 20}, {"S3", 300, 100}, {"S4", 400, 120}}
	mantenmientos := []Mantenimiento{{"M1", 10}, {"M2", 20}, {"M3", 30}, {"M4", 40}, {"M5", 50}, {"M1", 10}, {"M2", 20}, {"M3", 30}, {"M4", 40}, {"M5", 50}, {"M1", 10}, {"M2", 20}, {"M3", 30}, {"M4", 40}, {"M5", 50}, {"M1", 10}, {"M2", 20}, {"M3", 30}, {"M4", 40}, {"M5", 50}, {"M1", 10}, {"M2", 20}, {"M3", 30}, {"M4", 40}, {"M5", 50}, {"M1", 10}, {"M2", 20}, {"M3", 30}, {"M4", 40}, {"M5", 50}, {"M1", 10}, {"M2", 20}, {"M3", 30}, {"M4", 40}, {"M5", 50}}

	sumaProd := make(chan float64)
	sumaServ := make(chan float64)
	sumaMant := make(chan float64)

	go sumarProductos(productos, sumaProd)
	go sumarServicios(servicios, sumaServ)
	go sumarMantenimiento(mantenmientos, sumaMant)

	fmt.Printf("\nTotal suma mantenimientos: %.2f\n", <-sumaMant)

	fmt.Printf("\nTotal suma servicios: %.2f\n", <-sumaServ)

	fmt.Printf("\nTotal suma productos: %.2f\n", <-sumaProd)

}

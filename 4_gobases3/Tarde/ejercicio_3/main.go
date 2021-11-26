package main

import (
	"fmt"
	"math"
)

type Producto struct {
	nombre   string
	precio   float64
	cantidad int
}
type Servicio struct {
	nombre        string
	precio        float64
	mintrabajados int
}
type Mantenimiento struct {
	nombre string
	precio float64
}

func sumPro(p []Producto) float64 {
	total := 0.0
	for _, producto := range p {
		total += (producto.precio * float64(producto.cantidad))
	}
	return total
}

func sumSer(s []Servicio) float64 {
	total := 0.0
	for _, servicio := range s {
		total += servicio.precio * float64(Roundf(float64(servicio.mintrabajados)/30))
	}

	return total
}
func Roundf(x float64) int {
	t := math.Trunc(x)
	if math.Abs(x-t) > 0 {
		return int(t + math.Copysign(1, x))
	}
	return int(t)
}

func sumMant(m []Mantenimiento) float64 {
	total := 0.0
	for _, mant := range m {
		total += mant.precio
	}
	return total
}
func proceso1(i int, c chan int) {
	pr1 := Producto{
		nombre:   "mouse",
		precio:   5000.00,
		cantidad: 3,
	}
	pr2 := Producto{
		nombre:   "teclado",
		precio:   1000.00,
		cantidad: 3,
	}
	var listaProd []Producto
	listaProd = append(listaProd, pr1)
	listaProd = append(listaProd, pr2)
	fmt.Println("Suma productos:", sumPro((listaProd)))
	c <- i
}

func proceso2(i int, c chan int) {
	sr1 := Servicio{
		nombre:        "instalaciones",
		precio:        2000.00,
		mintrabajados: 30,
	}
	sr2 := Servicio{
		nombre:        "correciones",
		precio:        500.00,
		mintrabajados: 35,
	}
	var listaSer []Servicio
	listaSer = append(listaSer, sr1)
	listaSer = append(listaSer, sr2)
	fmt.Println("Suma servicios: ", sumSer((listaSer)))
	c <- i
}

func proceso3(i int, c chan int) {
	mn1 := Mantenimiento{
		nombre: "preventivo",
		precio: 300.00,
	}
	mn2 := Mantenimiento{
		nombre: "correctivo",
		precio: 800.00,
	}
	var listaMan []Mantenimiento
	listaMan = append(listaMan, mn1)
	listaMan = append(listaMan, mn2)
	fmt.Println("Suma Mantenimientos: ", sumMant((listaMan)))
	c <- i
}

func main() {
	c := make(chan int)
	// for i := 0; i < 1; i++ {
	go proceso1(0, c)
	go proceso2(1, c)
	go proceso3(2, c)
	// }
	for i := 0; i < 3; i++ {
		fmt.Println("termino : ", <-c)

	}

}

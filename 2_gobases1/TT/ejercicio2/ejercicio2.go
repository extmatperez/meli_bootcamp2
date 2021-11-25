package main

import "fmt"

func main() {
	var precio float32 //= 100
	const descuento float32 = 25

	fmt.Printf("Ingrese el valor del producto:")
	fmt.Scanf("%f", &precio)

	var a_pagar float32 = precio - descuento*precio/100
	fmt.Printf("El precio a pagar es: %0.2f\n", a_pagar)
}

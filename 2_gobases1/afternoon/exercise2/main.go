package main

import "fmt"

func main() {
	/*
		Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
		1 - Crear la aplicación de acuerdo a los requerimientos.
	*/
	price := 100.00
	discount := 40.00

	fmt.Printf("\nEl precio con descuento es: %.2f\n", price*0.01*(100-discount))
}

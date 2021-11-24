package main

import "fmt"

/*
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos,
para ello necesitan una aplicación que les permita calcular el descuento en base
a 2 variables, su precio y el descuento en porcentaje. Espera obtener como
resultado el valor con el descuento aplicado y luego imprimirlo en consola.
	- Crear la aplicación de acuerdo a los requerimientos.
*/

func main() {
	precio := 35000
	descuento := 50

	fmt.Printf("El valor de la ropa es: %d", precio)
	fmt.Printf("\nEl descuento es de un: %d%%", descuento)

	valor_desc := ((precio * 100) / descuento)
	fmt.Printf("\nEl valor de la ropa con descuento es: %d", valor_desc)
}

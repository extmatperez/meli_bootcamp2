// Ejercicio 2 - Descuento

// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
// Crear la aplicación de acuerdo a los requerimientos.

package main

import "fmt"

func main() {
	var price float32
	var porcent int

	price = 200.5
	porcent = 5

	if price > 0 && porcent > 0 && porcent <= 100 {
		discount := price * float32(porcent) / 100
		fmt.Printf("\nPrecio con descuento: %v", price-discount)
	} else {
		fmt.Println("\nError al aplicar descuento")
	}
}

/*
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello necesitan una
aplicación que les permita calcular el descuento en base a 2 variables, su precio y el descuento en porcentaje.
Espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.
*/

package main

import "fmt"

func main() {

	precio := 156.5
	descuento := 20.0

	precioFinal := precio - (descuento * precio) / 100

	fmt.Println("El valor con el descuento aplicado es $", precioFinal)
}
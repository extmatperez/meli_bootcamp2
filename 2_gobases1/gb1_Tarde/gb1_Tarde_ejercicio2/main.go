package main

import "fmt"

// Ejercicio 2 - Descuento
// Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola. 

func main () {

	price := 150.00
	discount := 30.00

	price_with_discount := price - (price*discount/100)

	fmt.Printf("The special price is: $ %.2f\n", price_with_discount)


}
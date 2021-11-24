/*Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos, para ello necesitan una aplicación que les permita calcular el descuento en base a 2 variables, su precio y el descuento en porcentaje. Espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.
*/
package main

import (
	"fmt"
)

func main() {

	var precio, descuento, valor float64
	fmt.Printf("Ingrese el precio del producto:")
	fmt.Scanf("%f", &precio)
	fmt.Printf("Ingrese el descuento del producto:")
	fmt.Scanf("%f", &descuento)

	valor = precio - ((descuento * precio) / 100)

	fmt.Printf("El valor del producto con descuento es de: %.2f", valor)

}

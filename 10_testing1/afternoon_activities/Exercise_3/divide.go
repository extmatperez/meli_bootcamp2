/* Ejercicio 3 - Test Unitario Método Dividir
Para el Método Dividir, visto en la clase:

// Función que recibe dos enteros (numerador, denominador) y retorna la división resultante
func Dividir(num, den int) int {
	return num / den
}

Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación en la que si el denominador es igual a 0,  retornar un error cuyo mensaje sea “El denominador no puede ser 0”. Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.
Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo dividir_test.go con el test diseñado.
*/

package divide

import "errors"

func Divide(num, den float64) (float64, error) {
	if num == 0 {
		return 0, errors.New("You can't divide a number for zero.")
	}
	return num / den, nil
}

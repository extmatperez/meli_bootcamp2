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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divide(t *testing.T) {
	num := 150.0
	den := 30.0
	expect := 5.0

	test, err := Divide(num, den)

	if err != nil {
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), "You can't divide a number for zero.")
		}
	}
	assert.Nil(t, err)
	assert.Equal(t, expect, test, "Something went wrong.")
}

/*
Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación
en la que si el denominador es igual a 0,  retornar un error cuyo mensaje sea “El denominador no puede ser 0”.
Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.
Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo dividir_test.go con el test diseñado.
*/

package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {

	//Arrenge
	num1 := 3
	num2 := 3
	resultadoEsperado := 1

	//Act
	resultado, err := Dividir(num1, num2)

	//Assert
	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, resultadoEsperado, resultado, "El resultado no fue el esperado")
	}
}

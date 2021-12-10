/*
Para el método Restar() visto en la clase, realizar el test unitario correspondiente. Para esto:
Dentro de la carpeta go-testing crear un archivo calculadora.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo calculadora_test.go con el test diseñado.
*/

package calculadoraTT

import (
	"testing"

	//"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {

	//Arrenge
	num1 := 3
	num2 := 10
	resultadoEsperado := -5

	//Act
	resultado := Restar(num1, num2)

	//Assert
	if resultadoEsperado != resultado {
		t.Errorf("El resultado de la función restar fue: %v, pero el resultado esperado era: %v", resultado, resultadoEsperado)
	}
}

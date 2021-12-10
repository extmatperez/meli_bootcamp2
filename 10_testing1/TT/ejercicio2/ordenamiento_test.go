/*
Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente, posteriormente diseñar un test unitario que valide el funcionamiento del mismo.
Dentro de la carpeta go-testing crear un archivo ordenamiento.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo ordenamiento_test.go con el test diseñado.
*/

package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {

	//Arrenge
	slice := []int{400, 600, 100, 300, 500, 200, 900}
	resultadoEsperado := []int{100, 200, 300, 400, 500, 600, 900}

	//Act
	resultado := Ordenar(slice)

	//Assert
	assert.Equal(t, resultadoEsperado, resultado, "El test falló")
}

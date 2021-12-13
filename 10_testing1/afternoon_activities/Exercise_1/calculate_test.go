/* Ejercicio 1 - Test Unitario Restar
Para el método Restar() visto en la clase, realizar el test unitario correspondiente. Para esto:
Dentro de la carpeta go-testing crear un archivo calculadora.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo calculadora_test.go con el test diseñado.
*/

package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subtract(t *testing.T) {
	// Datos a ser utilizados en el test
	num_1 := 5.0
	num_2 := 4.0
	result := 2.0

	// Se ejecuta el test
	expected := Subtract(num_1, num_2)

	// Validamos los resultados
	assert.Equal(t, result, expected, "There should be equals")
}

/* Diseñar un método que reciba un slice de enteros y los ordene de forma ascendente, posteriormente diseñar un test unitario que
valide el funcionamiento del mismo.
Dentro de la carpeta go-testing crear un archivo ordenamiento.go con la función a probar.
Dentro de la carpeta go-testing crear un archivo ordenamiento_test.go con el test diseñado.
*/

package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_order(t *testing.T) {
	slice_int := []int{3, 5, 7, 1, 4, 2, 6}

	Order(slice_int)
	// Example to fail test
	slice_order := []int{1, 2, 3, 4, 5, 7, 6}

	assert.Equal(t, slice_int, slice_order, "The slice is unorderer.")
}

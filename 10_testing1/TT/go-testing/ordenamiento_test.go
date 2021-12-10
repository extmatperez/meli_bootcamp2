package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	sliceToOrder := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	sliceResponse := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sliceToOrder = Ordenar(sliceToOrder)

	assert.Equal(t, len(sliceResponse), len(sliceToOrder), "Los slices deben tener el mismo tamaño")

	for i, v := range sliceToOrder {
		assert.Equal(t, v, sliceResponse[i], "Deben ser iguales")
	}
}

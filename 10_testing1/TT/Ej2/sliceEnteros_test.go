package slice_enteros_test

import (
	"testing"

	slice_enteros "github.com/extmatperez/meli_bootcamp2/10_testing1/TT/Ej2"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	arg := []int{3, 6, 2, 5, 1, 4}

	s := slice_enteros.Slice(arg)
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6}

	assert.Equal(t, resultadoEsperado, s, "Deberian ser iguales")
}

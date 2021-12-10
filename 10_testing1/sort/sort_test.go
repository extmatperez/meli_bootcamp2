package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	sliceNum := []int{6, 5, 7, 2, 4, 12, 9, 75, 3}
	resultadoEsperado := []int{2, 3, 4, 5, 6, 7, 9, 12, 75}

	resultado := MethodSort(sliceNum)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

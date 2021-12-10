package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {

	slice := []int{3, 1, 4, 2, 5}
	sliceEsperado := []int{1, 2, 3, 4, 6}

	result := Ordernar(slice)

	assert.Equal(t, sliceEsperado, result, "slice no esta correctamente ordenado")
}

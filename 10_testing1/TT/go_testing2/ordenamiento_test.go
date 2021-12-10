package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	slice := []int{1, 5, 7, 9, 3, 2, 4, 8, 6}
	resultadoEsperado := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	resultado := Ordenar(slice)
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")

}

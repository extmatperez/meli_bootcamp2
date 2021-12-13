package ordenamiento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	lista := []int{2, 5, 6, 4, 8, 9, 1}
	resultadoEsperado := []int{1, 2, 4, 5, 6, 8, 9}
	resultado := Ordenar(lista)
	assert.Equal(t, resultadoEsperado, resultado, lista)
}

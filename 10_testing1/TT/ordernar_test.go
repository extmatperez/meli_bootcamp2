package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {

	lista := []int{4, 2, 5, 1, 6, 7}

	resultadoEsperado := []int{1, 2, 4, 5, 6, 7}

	resultado := Ordenar(lista)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

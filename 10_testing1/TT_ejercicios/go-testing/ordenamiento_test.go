package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenar(t *testing.T) {
	desordenado := []int{90, 8, 4, 10, 3, 2}
	resultadoEsperado := []int{2, 3, 4, 8, 10, 90}

	resultado := Ordenar(desordenado)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

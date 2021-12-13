package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num1 := 5
	num2 := 5
	resultadoEsperado := 1

	resultado, err := Dividir(num1, num2)

	assert.Nil(t, nil, err)
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}

func TestDividirBad(t *testing.T) {
	num1 := 5
	num2 := 0
	resultadoEsperado := 0

	resultado, err := Dividir(num1, num2)

	assert.Error(t, err, "Hubo un error")
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}

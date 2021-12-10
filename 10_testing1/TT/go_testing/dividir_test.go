package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	num1 := 3
	num2 := 1
	// num3 := 0

	resultadoEsperado := 3

	resultado1, err := Divide(num1, num2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resultadoEsperado, resultado1, "Deben ser iguales")

	resultado2, err := Divide(num1, num2)
	if err != nil {
		t.Error(err)
	}
	// assert.Nil(t, resultado2)
	assert.Equal(t, resultadoEsperado, resultado2, "Deben ser iguales")

}

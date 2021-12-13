package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	num1 := 5
	num2 := 5
	resultadoEsperado := 10

	resultado := Sumar(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("Error al sumar")
	}
}

func TestRestar(t *testing.T) {
	num1 := 5
	num2 := 5
	resultadoEsperado := 0

	resultado := Resta(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}

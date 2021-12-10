package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	num1 := 100
	num2 := 200
	resultadoEsperado := -300

	resultado := Sumar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "Resultado suma")
}

func TestRestar(t *testing.T) {
	num1 := 100
	num2 := 200
	resultadoEsperado := -200

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "Resultado resta")
}

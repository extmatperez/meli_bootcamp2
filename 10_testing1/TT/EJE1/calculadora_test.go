package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 10
	num2 := 5
	resultadoEsperado := 9
	resultado := Restar(num1, num2)
	assert.Equal(t, resultadoEsperado, resultado, "Deben ser iguales")
}

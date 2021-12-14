package TT

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	num1 := 100
	num2 := 10
	resultadoEsperado := 90

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "debe dar 90")
}

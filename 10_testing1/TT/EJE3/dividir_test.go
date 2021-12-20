package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividirSuccess(t *testing.T) {
	num1 := 10
	num2 := 5
	resultadoEsperado := 2
	resultado, _ := Dividir(num1, num2)
	assert.Equal(t, resultadoEsperado, resultado, "La division no da el resultado esperado")
}

func TestDividirError(t *testing.T) {
	num1 := 10
	num2 := 0
	_, err := Dividir(num1, num2)
	assert.NotNil(t, err)
}

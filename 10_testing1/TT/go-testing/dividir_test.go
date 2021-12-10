package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 4
	num2 := 2
	resultadoEsperado := 2
	// Se ejecuta el test
	resultado, err := Dividir(num1, num2)

	// Se validan los resultados aprovechando testify
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado, "El resultado debe ser igual")

}

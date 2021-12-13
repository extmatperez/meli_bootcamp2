package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestDividir1(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 6.0
	num2 := 3.0
	resultadoEsperado := 2.0
	// Se ejecuta el test
	resultado, err := Dividir(num1, num2)
	// Se validan los resultados
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividir2(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 6.0
	num2 := 0.0
	// Se ejecuta el test
	_, err := Dividir(num1, num2)
	// Se validan los resultados
	assert.NotNil(t, err)
}

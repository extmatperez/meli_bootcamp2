package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	resultadoEsperado := 8
	// Se ejecuta el test
	resultado := Sumar(num1, num2)
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 8
	num2 := 5
	resultadoEsperado := 3
	// Se ejecuta el test
	resultado := Restar(num1, num2)
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

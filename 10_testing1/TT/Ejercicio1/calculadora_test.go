package calculadora

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 10
	num2 := 5
	resultadoEsperado := 5
	// Se ejecuta el test
	resultado := Restar(num1, num2)
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "Funcion restar() arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
}

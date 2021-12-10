package ordenamiento

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {

	slice := Nums{
		num1: 4,
		num2: 3,
		num3: 2,
		num4: 1,
	}
	var resultadoEsperado []int
	resultadoEsperado = append(resultadoEsperado, slice.num4, slice.num3, slice.num2, slice.num1)

	resultado := Nums.Order(slice)
	// Se validan los resultados
	assert.Equal(t, resultadoEsperado, resultado, "Funcion Order() arrojo el resultado = %v, pero el esperado es %v", resultado, resultadoEsperado)
}

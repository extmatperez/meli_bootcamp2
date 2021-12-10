package calculadora

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumar(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	// Este es el ARRANGE.
	n1 := 3
	n2 := 5
	promisedValue := 8

	// Se ejecuta el test. Este es el ACT.
	res := Sumar(n1, n2)

	// Se validan los resultados. Este es el ASSERT.
	// Usar assert es lo mismo que validar la correspondencia y responder con el resultado del test.
	/*if res != promisedValue {
		t.Errorf("Funcion Suma() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Suma() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
}

func TestRestar(t *testing.T) {
	n1 := 10
	n2 := 6
	promisedValue := 4

	res := Restar(n1, n2)

	/*if res != promisedValue {
		t.Errorf("Funcion Resta() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Resta() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
}

func TestMultiplicar(t *testing.T) {
	n1 := 8
	n2 := 5
	promisedValue := 40

	res := Multiplicar(n1, n2)

	/*if res != promisedValue {
		t.Errorf("Funcion Multiplicar() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Multiplicar() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
}

func TestDividir(t *testing.T) {
	n1 := 5
	n2 := 2
	promisedValue := 2

	res, err := Dividir(n1, n2)

	/*if res != promisedValue {
		t.Errorf("Funcion Dividir() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, promisedValue, res, fmt.Sprintf("Funcion Dividir() arrojo el resultado %v, pero el esperado es %v", res, promisedValue))
	}
}

func TestOrdenarAsc(t *testing.T) {
	nums := []int{9, 12, 3, 2, 6, 5}
	promised := []int{2, 3, 5, 6, 9, 12}

	res := OrdenarAsc(nums)

	/*if res != promisedValue {
		t.Errorf("Funcion Dividir() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promised, res, fmt.Sprintf("Funcion OrdenarAsc() arrojo el resultado %v, pero el esperado es %v", res, promised))
}

func TestOrdenarDesc(t *testing.T) {
	nums := []int{9, 12, 3, 2, 6, 5}
	promised := []int{12, 9, 6, 5, 3, 2}

	res := OrdenarDesc(nums)

	/*if res != promisedValue {
		t.Errorf("Funcion Dividir() arrojo el resultado = %v, pero el esperado es %v", res, promisedValue)
	}*/
	assert.Equal(t, promised, res, fmt.Sprintf("Funcion OrdenarDesc() arrojo el resultado %v, pero el esperado es %v", res, promised))
}

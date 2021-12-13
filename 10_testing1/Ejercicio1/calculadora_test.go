package calculadora

import "testing"

func TestSumar(t *testing.T) {

	num1 := 1
	num2 := 2
	resEsperado := 3

	result := Sumar(num1, num2)

	if result != resEsperado {
		t.Errorf("funcion suma() arrojo el resulado %v, pero el resultado esperado es %v", result, resEsperado)
	}
}
func TestRestar(t *testing.T) {

	num1 := 1
	num2 := 2
	resEsperado := -1

	result := Restar(num1, num2)

	if result != resEsperado {
		t.Errorf("funcion Restar() arrojo el resulado %v, pero el resultado esperado es %v", result, resEsperado)
	}
}

package calculadora

import "testing"


func TestRestar(t *testing.T) {
	num1 := 5
	num2 := 3
	resultadoEsperado := 2

	resultado := Restar(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("La funcion suma() arrojo el resultado = %v, pero el estado es %v", resultado, resultadoEsperado)
	}
}
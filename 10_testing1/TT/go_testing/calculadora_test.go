package calculadora

import "testing"

func TestRestar(t *testing.T) {
	num1 := 10
	num2 := 5
	resultadoEsperado := 5

	resultado := Restar(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("Funcion Suma() arrojo el resultado de = %v, pero esperaba que sea = %v", resultado, resultadoEsperado)
	}
}

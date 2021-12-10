package calculadora

import "testing"

func TestSumar(t *testing.T) {

	n1, n2 := 3, 4
	resultadoEsperado := 7

	resultado := Sumar(n1, n2)

	if resultado != resultadoEsperado {
		t.Errorf("function TestSumar FAILED, expect %v, got %v", resultadoEsperado, resultado)
	}

}

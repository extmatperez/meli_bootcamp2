package funciones

import "testing"

func TestRestar(t *testing.T) {
	num1 := 10
	num2 := 5
	resultadoEsperado := 5

	resultado := Restar(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("Funcion Restar() arroja el resultado %v, pero espera el valor %v", resultado, resultadoEsperado)
	}
}

func TestRestar2(t *testing.T) {
	num1 := 30
	num2 := 40
	resultadoEsperado := -10

	resultado := Restar(num1, num2)

	if resultado != resultadoEsperado {
		t.Errorf("Funcion Restar() arroja el resultado %v, pero espera el valor %v", resultado, resultadoEsperado)
	}

}

func TestOrdenar(t *testing.T) {
	slice1 := []int{9, 5, 3, 7, 8, 1, 6, 9}
	resultadoEsperado := []int{1, 3, 5, 6, 7, 8, 9, 9}

	resultado := Ordenar(slice1)

	if len(slice1) != len(resultadoEsperado) {
		t.Errorf("Cambia el largo del Slice")
		return
	}

	for i, _ := range slice1 {
		if resultado[i] != resultadoEsperado[i] {
			t.Errorf("Organizo todo mal")
			return
		}
	}
}

func TestOrdenar2(t *testing.T) {
	slice1 := []int{1, 3, 5, 6, 7, 8, 9, 9}
	resultadoEsperado := []int{1, 3, 5, 6, 7, 8, 9, 9}

	resultado := Ordenar(slice1)

	if len(slice1) != len(resultadoEsperado) {
		t.Errorf("Cambia el largo del Slice")
		return
	}

	for i, _ := range slice1 {
		if resultado[i] != resultadoEsperado[i] {
			t.Errorf("No mantiene el ordenado")
			return
		}
	}
}

func TestDividir(t *testing.T) {
	num1 := 20
	num2 := 2
	resultadoEsperado := 10

	resultado, errDividir := Dividir(num1, num2)

	if errDividir != nil {
		t.Error(errDividir)
		return
	}

	if resultado != float64(resultadoEsperado) {
		t.Errorf("Funcion Dividir() arroja el resultado %v, pero espera el valor %v", resultado, resultadoEsperado)
		return
	}
}

func TestDividir2(t *testing.T) {
	num1 := 31
	num2 := 2
	resultadoEsperado := 15.5

	resultado, errDividir := Dividir(num1, num2)

	if errDividir != nil {
		t.Error(errDividir)
		return
	}

	if resultado != float64(resultadoEsperado) {
		t.Errorf("Funcion Dividir() arroja el resultado %v, pero espera el valor %v", resultado, resultadoEsperado)
		return
	}
}

func TestDividir3(t *testing.T) {

	num1 := 20
	num2 := 0
	resultadoEsperado := 10

	resultado, errDividir := Dividir(num1, num2)

	if errDividir != nil {
		t.Error(errDividir)
		return
	}

	if resultado != float64(resultadoEsperado) {
		t.Errorf("Funcion Dividir() arroja el resultado %v, pero espera el valor %v", resultado, resultadoEsperado)
		return
	}
}

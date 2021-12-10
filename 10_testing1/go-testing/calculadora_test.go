package go_testing

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumar(t *testing.T) {
	num1 := 5
	num2 := 50
	expectedResult := -45

	result := Restar(num1, num2)

	if result != expectedResult {
		t.Errorf("Funcion Restar(num1, num2, int) int => Arrojo el resultado %v, pero el resultado esperado es %v", result, expectedResult)
	}
}

func TestDividir(t *testing.T){
	num1 := 33.0
	num2 := 3.0
	expectedResult := 141.0

	result, err := Dividir(num1, num2)

	fmt.Println(result)
	fmt.Println(expectedResult)

	if err != nil {
		t.Errorf("El denominador no puede ser 0")
		return
	}
	if result == expectedResult {
		assert.Equal(t, expectedResult, result, "deben ser iguales")
	} else {
		t.Errorf("Funcion Dividir(num1, num2, float32) (float32, error) => Arrojo el resultado %v, pero el resultado esperado es %v", result, expectedResult)
	}
	/*if result != expectedResult{
		t.Errorf("Funcion Dividir(num1, num2, float32) (float32, error) => Arrojo el resultado %v, pero el resultado esperado es %v", result, expectedResult)
	}*/
}
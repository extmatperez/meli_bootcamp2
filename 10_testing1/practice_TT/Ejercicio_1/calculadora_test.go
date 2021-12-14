package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSubstract(t *testing.T) {
	var (
		n1             = 5
		n2             = 3
		expectedResult = 2
	)

	result := Restar(n1, n2)

	if result != expectedResult {
		t.Errorf("The function subtract obtained %v but the expected result was %v", result, expectedResult)
	}
}
func TestSplit(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 1

	//Act
	resultado := Dividir(num1, num2)

	//Assert
	assert.NotNil(t, resultado)
}

package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstractEqual(t *testing.T) {
	// (input/output)
	num1 := 2
	num2 := 2
	expectedResult := 0

	// Execute test
	result := Substract(num1, num2)

	//check
	assert.Equal(t, expectedResult, result, "they should be equal")
}

func TestSubstractNotEqual(t *testing.T) {
	// (input/output)
	num1 := 2
	num2 := 0
	expectedResult := 0

	// Execute test
	result := Substract(num1, num2)

	//check
	assert.NotEqual(t, expectedResult, result, "they should be not equal")
}

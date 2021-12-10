package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideSuccess(t *testing.T) {
	// (input/output)
	num1 := 2
	num2 := 1
	expectedResult := 2

	// Execute test
	result, _ := Divide(num1, num2)

	//check
	assert.Equal(t, expectedResult, result, "should be equal")
}

func TestDivideError(t *testing.T) {
	// (input/output)
	num1 := 2
	num2 := 0

	// Execute test
	_, err := Divide(num1, num2)

	//check
	assert.NotNil(t, err)
}

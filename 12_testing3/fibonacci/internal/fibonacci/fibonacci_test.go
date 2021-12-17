package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	///Arrange
	var real int
	//Act
	result := Fibonacci()
	for i := 0; i < 5; i++ {
		real = result(i)
	}
	//Assert
	assert.Equal(t, 3, real, "Should be equals")
}

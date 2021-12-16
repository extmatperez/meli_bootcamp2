package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {

	///Arrange

	//Act
	result := Fibonacci()
	fmt.Println(result)

	//Assert
	assert.Equal(t, result, Fibonacci())

}

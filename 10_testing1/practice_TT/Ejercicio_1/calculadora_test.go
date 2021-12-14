package calculadora

import "testing"

func testRestar(t *testing.T) {
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

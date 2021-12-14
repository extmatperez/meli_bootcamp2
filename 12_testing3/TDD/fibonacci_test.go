package fibonacci

import "testing"

func TestFibonacciRecursiva(t *testing.T) {
	tests := []struct {
		arg      int
		expected int
	}{{0, 0}, {1, 1}, {2, 1}, {8, 21}, {40, 102334155}}
	for i, d := range tests {
		got := fibonacciRecursiva(d.arg)
		if got != d.expected {
			t.Errorf("Test[%d]: fibonacciRecursiva(%d) returned %d instead of %d", i, d.arg, got, d.expected)
		}
	}
}

func TestFibonacciIterativa(t *testing.T) {
	tests := []struct {
		arg      int
		expected int
	}{{0, 0}, {1, 1}, {2, 1}, {8, 21}, {40, 102334155}}
	for i, d := range tests {
		got := fibonacciRecursiva(d.arg)
		if got != d.expected {
			t.Errorf("Test[%d]: fibonacciIterativa(%d) returned %d instead of %d", i, d.arg, got, d.expected)
		}
	}
}

func BenchmarkFibonacciRecursiva(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fibonacciRecursiva(40)
	}
}

func BenchmarkFibonacciIterativa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciIterativa(40)
	}
}

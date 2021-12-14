package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{8, 21},
	}

	for i, test := range tests {
		result := fibonacci(test.in)

		if result != test.out {
			t.Errorf("Test %d FAIL", i+1)
		}
	}
}

func TestFibonacciMostEfficient(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{8, 21},
	}

	for i, test := range tests {
		result := fibonacciMostEfficient(test.in)

		if result != test.out {
			t.Errorf("Test %d FAIL", i+1)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	tests := []struct {
		in  int
		out int
	}{
		{0, 0},
		{8, 21},
		{20, 6765},
		{30, 832040},
		{36, 14930352},
	}

	for _, test := range tests {
		_ = fibonacci(test.in)
	}
}

func BenchmarkFibonacciMostEfficient(b *testing.B) {
	tests := []struct {
		in  int
		out int
	}{
		{0, 0},
		{8, 21},
		{20, 6765},
		{30, 832040},
		{36, 14930352},
	}

	for _, test := range tests {
		_ = fibonacci(test.in)
	}
}

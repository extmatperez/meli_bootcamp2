package fibonacci

func fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	}

	return fibonacci(number-1) + fibonacci(number-2)
}

func fibonacciMostEfficient(number int) int {
	sequence := []int{0, 1}

	for i := 2; i <= number; i++ {
		if len(sequence) > i {
			sequence[i] = sequence[i-2] + sequence[i-1]
		} else {
			sequence = append(sequence, sequence[i-2]+sequence[i-1])
		}
	}

	return sequence[number]
}

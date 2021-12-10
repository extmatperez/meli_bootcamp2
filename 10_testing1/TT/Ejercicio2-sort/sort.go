package sort

func Sort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		auxVal := numbers[i]
		for j := i - 1; j >= 0 && numbers[j] > auxVal; j-- {
			numbers[j+1] = numbers[j]
			numbers[j] = auxVal
		}
	}

	return numbers
}

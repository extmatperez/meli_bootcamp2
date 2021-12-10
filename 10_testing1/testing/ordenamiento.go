package testing

func SortByInsertion(numbers []int) {

	for i := 1; i < len(numbers); i++ {
		actual := numbers[i]
		j := i - 1
		for (j >= 0) && (numbers[j] > actual) {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = actual
	}
}

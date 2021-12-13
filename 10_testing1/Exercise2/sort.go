package ordenar

func Sort(list []int) []int {
	for i := 1; i < len(list); i++ {
		aux := list[i]
		for j := i - 1; j >= 0 && list[j] > aux; j-- {
			list[j+1] = list[j]
			list[j] = aux
		}
	}
	return list
}

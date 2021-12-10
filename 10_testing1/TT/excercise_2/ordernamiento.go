package ordenamiento

func Ordernar(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice)-1; j++ {
			if slice[i] > slice[j] {
				aux := slice[i]
				slice[i] = slice[j]
				slice[j] = aux
			}
		}
	}
	return slice
}

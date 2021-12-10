package testing

func OrdenamientoSliceAscendente(slice []int) []int {
	for i := range slice {
		for j := range slice {
			if slice[i] < slice[j] {
				aux := slice[i]
				slice[i] = slice[j]
				slice[j] = aux
			}
		}
	}
	return slice
}

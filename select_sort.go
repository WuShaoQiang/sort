package chapter7

func SelectSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		var min = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
	return a
}

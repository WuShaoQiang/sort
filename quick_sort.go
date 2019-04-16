package chapter7

func QuickSort(a []int) []int {
	quickSort(a, 0, len(a)-1)
	return a
}

func quickSort(a []int, left, right int) {
	if left < right {
		low := left
		high := right
		tmp := a[low]
		var leftSide = false

		for low < high {
			if !leftSide {
				if a[high] < tmp {
					a[low] = a[high]
					low++
					leftSide = true
				} else {
					high--
				}
			} else {
				if a[low] > tmp {
					a[high] = a[low]
					high--
					leftSide = false
				} else {
					low++
				}
			}
		}
		a[low] = tmp

		quickSort(a, left, low-1)
		quickSort(a, low+1, right)
	}

}

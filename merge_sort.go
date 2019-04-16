package chapter7

func MergeSort(a []int) []int {
	tmp := make([]int, len(a))
	sort(a, 0, len(a)-1, tmp)
	return a

}

func sort(a []int, left, right int, tmp []int) {
	if left < right {
		mid := (right + left) / 2
		sort(a, left, mid, tmp)
		sort(a, mid+1, right, tmp)
		merge(a, left, mid, right, tmp)
	}
}

func merge(a []int, left, mid, right int, tmp []int) {
	l := left
	r := mid + 1
	t := 0

	for l <= mid && r <= right {
		if a[l] <= a[r] {
			tmp[t] = a[l]
			t++
			l++
		} else {
			tmp[t] = a[r]
			t++
			r++
		}
	}

	for l <= mid {
		tmp[t] = a[l]
		t++
		l++
	}
	for r <= right {
		tmp[t] = a[r]
		t++
		r++
	}

	t = 0
	for left <= right {
		a[left] = tmp[t]
		left++
		t++
	}
}

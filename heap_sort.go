package chapter7

import "fmt"

func HeapSortUp(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		adjustHeapUp(a, i, len(a))
	}

	for j := len(a) - 1; j > 0; j-- {
		a[0], a[j] = a[j], a[0]
		adjustHeapUp(a, 0, j)
	}
	return a
}

func adjustHeapUp(a []int, i int, length int) {
	tmp := a[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && a[k+1] > a[k] {
			k++
		}

		if a[k] > tmp {
			a[i] = a[k]
			i = k
		} else {
			break
		}
	}
	a[i] = tmp
}

func HeapSortDown(a []int) []int {
	for k := len(a)/2 - 1; k >= 0; k-- {
		adjustHeapDown(a, k, len(a))
	}
	fmt.Println(a)

	for k := len(a) - 1; k > 0; k-- {
		a[0], a[k] = a[k], a[0]
		adjustHeapDown(a, 0, k)
	}
	return a
}

func adjustHeapDown(a []int, i int, length int) {
	tmp := a[i]
	for k := 2*i + 1; k < length; k = k*2 + 1 {
		if k+1 < length && a[k+1] < a[k] {
			k++
		}

		if a[k] < tmp {
			a[i] = a[k]
			i = k
		} else {
			break
		}
	}
	a[i] = tmp
}

package benchtest

import (
	"testing"

	"github.com/WuShaoQiang/data_structure_and_algorithm_analysis/chapter7"
)

func gerArray() []int {
	a := make([]int, 10000)
	for i := 10000; i > 0; i-- {
		a = append(a, i)
	}
	return a
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := gerArray()
		chapter7.BubbleSort(a)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := gerArray()

		chapter7.HeapSortUp(a)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := gerArray()

		chapter7.InsertSortUp(a)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := gerArray()

		chapter7.MergeSort(a)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := gerArray()

		chapter7.QuickSort(a)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := gerArray()

		chapter7.SelectSort(a)
	}
}

func BenchmarkShellSort(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		a := gerArray()

		chapter7.ShellSort(a)
	}
}

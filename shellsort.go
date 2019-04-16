package chapter7

func ShellSort(a []int) []int {
	for increment := len(a) / 2; increment > 0; increment = increment / 2 {
		for i := increment; i < len(a); i++ {
			//插入排序
			temp := a[i]
			var j int
			for j = i; j >= increment && a[j-increment] > temp; j -= increment {
				a[j] = a[j-increment]
			}
			a[j] = temp
		}
	}
	return a
}

package chapter7

// 错在了没将a[i]先取出来，导致后面a[i]已经被改变了
// 错在了没有考虑极端情况，比如元素要移到第一位
// O(N^2)
// 如果插入排序的数组几乎被排序，这种方法会比较快
func InsertSortUp(a []int) []int {
	var j int
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		for j = i; j > 0 && tmp < a[j-1]; j-- {
			a[j] = a[j-1]
		}
		a[j] = tmp
	}
	return a
}

// 降序插入排序和升序的只差了一个符号
func InsertSortDown(a []int) []int {
	var j int
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		for j = i; j > 0 && tmp > a[j-1]; j-- {
			a[j] = a[j-1]
		}
		a[j] = tmp
	}
	return a
}

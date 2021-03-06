package sort

//归并排序
func MergeSort(array []int) []int {
	n := len(array)
	if n < 2 {
		return array
	}

	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}

//merge 对两个已排序好的数组，进行合并
func merge(left []int, right []int) []int {
	tmp := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	//直到其中一个数组已遍历完
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmp = append(tmp, left[i])
			i++
		} else {
			tmp = append(tmp, right[j])
			j++
		}
	}
	//另外一个数组直接加到后面
	tmp = append(tmp, left[i:]...)
	tmp = append(tmp, right[j:]...)
	return tmp
}

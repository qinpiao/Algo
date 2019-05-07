package sort


func MergeSort(s []int){
	if len(s) <= 1 {
		return
	}
	mergeSortC(s, 0, len(s)-1)
}

func mergeSortC(s []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	mergeSortC(s, start, mid)
	mergeSortC(s, mid+1, end)
	merge(s, start, mid, end)
}

func merge(s []int, start, mid, end int){
	i, j, k := start, mid+1, 0

	temp := make([]int, end-start+1)
	for ;i<= mid && j<= end; k++ {
		if s[i] < s[j] {
			temp[k] = s[i]
			i++
		} else {
			temp[k] = s[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		temp[k] = s[i]
		k++
	}
	for ; j <= end; j++ {
		temp[k] = s[j]
		k++
	}
	copy(s[start:end+1], temp)
}
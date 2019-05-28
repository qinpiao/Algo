package sort


func QuickSort(s []int) {
	if len(s) <= 1 {
		return
	}
	quickSort(s, 0, len(s)-1)
}

func quickSort(s []int, start, end int) {
	if start >= end {
		return
	}
	pivot := end
	for i := end-1; i >= start; i-- {
		if s[i] > s[pivot] {
			temp := s[i]
			j := i
			for ;j< pivot; j++ {
				s[j] = s[j+1]
			}
			s[j], pivot = temp, j-1
		}
	}
	mergeSortC(s, start, pivot-1)
	mergeSortC(s, pivot, end)
}
package sort

import (
	"testing"
	"fmt"
)

func TestMergeSort(t *testing.T) {
	s := []int{6,5,4,3,2,1}
	MergeSort(s)
	fmt.Println(s)
	h := []int{2}
	MergeSort(h)
	fmt.Println(h)
	z := []int{1,2,3,4,5,6}
	MergeSort(z)
	fmt.Println(z)
}
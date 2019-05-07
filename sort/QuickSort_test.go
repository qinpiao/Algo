package sort

import (
	"testing"
	"fmt"
)

func TestQuickSort(t *testing.T) {
	s := []int{6,5,5,4,3,2,1}
	QuickSort(s)
	fmt.Println(s)
	h := []int{2,3,1,4,2,1,5,6}
	QuickSort(h)
	fmt.Println(h)
	z := []int{2,3,3,4,5,6,1}
	QuickSort(z)
	fmt.Println(z)
}
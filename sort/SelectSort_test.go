package sort

import (
	"testing"
	"fmt"
)

func TestSelectSort(t *testing.T) {
	s := []int{6,5,4,3,2,1}
	SelectSort(s)
	fmt.Println(s)
	h := []int{2}
	SelectSort(h)
	fmt.Println(h)
	z := []int{1,2,3,4,5,6}
	SelectSort(z)
	fmt.Println(z)
}
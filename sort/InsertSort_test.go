package sort

import (
	"testing"
	"fmt"
)

func TestInsertSort(t *testing.T) {
	s := []int{6,5,4,3,2,1}
	InsertSort(s)
	fmt.Println(s)
	h := []int{2}
	InsertSort(h)
	fmt.Println(h)
	z := []int{1,2,3,4,5,6}
	InsertSort(z)
	fmt.Println(z)
}
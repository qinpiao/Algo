package sort

import (
	"testing"

	"fmt"
)

func TestBubleSort(t *testing.T) {
	s := []int{6,5,4,3,2,1}
	BubleSort(s)
	fmt.Println(s)
	h := []int{2}
	BubleSort(h)
	fmt.Println(h)
}

func TestBubleSort2(t *testing.T) {
	s := []int{6,5,4,3,2,1}
	BubleSort2(s)
	fmt.Println(s)
	h := []int{2}
	BubleSort2(h)
	fmt.Println(h)
}

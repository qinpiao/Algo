package heap

import (
	"testing"
	"fmt"
)

func TestTopMaxHeap(t *testing.T) {
	hp := NewTopMaxHeap(10)
	s := []int{2,4,3,1,5,6,7,8}
	for _, e := range s{
		if err := hp.Insert(e); err != nil {
			panic(err)
		}
	}
	hp.String()
	fmt.Println(hp.removeMax())
	hp.String()
	hp.Insert(2)
	hp.String()
	hp.Insert(8)
	hp.String()
}

func TestTopMinHeap(t *testing.T) {
	hp := NewTopMinHeap(10)
	s := []int{2,4,3,1,5,6,7,8}
	for _, e := range s{
		if err := hp.Insert(e); err != nil {
			panic(err)
		}
	}
	hp.String()
	fmt.Println(hp.removeMin())
	hp.String()
	hp.Insert(2)
	hp.String()
	hp.Insert(8)
	hp.String()
}
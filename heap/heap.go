package heap

import (
	"errors"
	"fmt"
)

type heap struct {
	cap int
	len int
	ls []int
}

type topMaxHeap struct {
	*heap
}

type topMinHeap struct {
	*heap
}

func newHeap(capacity int) *heap {
	newHeap := &heap{cap:capacity, len:0}
	newHeap.ls = make([]int, capacity+1)
	return newHeap
}

func NewTopMaxHeap(capacity int) *topMaxHeap {
	newHeap := &topMaxHeap{heap:newHeap(capacity)}
	return newHeap
}

func NewTopMinHeap(capacity int) *topMinHeap {
	newHeap := &topMinHeap{heap:newHeap(capacity)}
	return newHeap
}


func (hp *topMaxHeap) Insert(data int) error {
	if hp.len == hp.cap {
		return errors.New("Heap full!")
	}
	hp.len += 1
	hp.ls[hp.len] = data
	cIndex := hp.len
	pIndex := cIndex / 2
	for pIndex > 0 && hp.ls[cIndex] > hp.ls[pIndex] {
		hp.ls[cIndex], hp.ls[pIndex] = hp.ls[pIndex], hp.ls[cIndex]
		cIndex, pIndex = pIndex, pIndex/2
	}
	return nil
}

func (hp *topMinHeap) Insert(data int) error {
	if hp.len == hp.cap {
		return errors.New("Heap full!")
	}
	hp.len += 1
	hp.ls[hp.len] = data

	cIndex := hp.len
	pIndex := cIndex / 2
	for pIndex > 0 && hp.ls[cIndex] < hp.ls[pIndex] {
		hp.ls[cIndex], hp.ls[pIndex] = hp.ls[pIndex], hp.ls[cIndex]
		cIndex, pIndex = pIndex, pIndex/2
	}
	return nil
}

func (hp *topMaxHeap) removeMax() int {
	deleted := hp.ls[1]
	if hp.len == 1 {
		hp.ls = []int{}
		hp.len = 0
		return deleted
	}
	hp.ls[1] = hp.ls[hp.len]
	hp.len -= 1
	nextIndex := 1
	for nextIndex < hp.len {
		left := nextIndex * 2
		right := left + 1
		if left < hp.len && right < hp.len {
			if hp.ls[left] > hp.ls[right] {
				hp.ls[nextIndex], hp.ls[left] = hp.ls[left], hp.ls[nextIndex]
				nextIndex = left
			} else {
				hp.ls[nextIndex], hp.ls[right] = hp.ls[right], hp.ls[nextIndex]
				nextIndex = right
			}
		} else if left < hp.len && right > hp.len {
			hp.ls[nextIndex], hp.ls[left] = hp.ls[left], hp.ls[nextIndex]
			nextIndex = left
		} else {
			break
		}
	}
	return deleted
}

func (hp *topMinHeap) removeMin() int {
	deleted := hp.ls[1]
	if hp.len == 1 {
		hp.ls = []int{}
		hp.len = 0
		return deleted
	}
	hp.ls[1] = hp.ls[hp.len]
	hp.len -= 1
	nextIndex := 1
	for nextIndex < hp.len {
		left := nextIndex * 2
		right := left + 1
		if left < hp.len && right < hp.len {
			if hp.ls[left] < hp.ls[right] {
				hp.ls[nextIndex], hp.ls[left] = hp.ls[left], hp.ls[nextIndex]
				nextIndex = left
			} else {
				hp.ls[nextIndex], hp.ls[right] = hp.ls[right], hp.ls[nextIndex]
				nextIndex = right
			}
		} else if left < hp.len && right > hp.len {
			hp.ls[nextIndex], hp.ls[left] = hp.ls[left], hp.ls[nextIndex]
			nextIndex = left
		} else {
			break
		}
	}
	return deleted
}


func (hp *topMaxHeap) String() {
	for i := 1; i <= hp.len; i++ {
		fmt.Printf("%d ", hp.ls[i])
	}
	fmt.Println()
}

func (hp *topMinHeap) String() {
	for i := 1; i <= hp.len; i++ {
		fmt.Printf("%d ", hp.ls[i])
	}
	fmt.Println()
}

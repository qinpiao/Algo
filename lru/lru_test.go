package lru

import (
	"testing"

	"fmt"
)

func TestLruCache(t *testing.T) {
	lru := NewLruCache(5)
	if err := lru.Put("1", 1); err != nil {
		panic(err)
	}
	if err := lru.Put("2", 2); err != nil {
		panic(err)
	}
	if err := lru.Put("3", 3); err != nil {
		panic(err)
	}
	if err := lru.Put("4", 4); err != nil {
		panic(err)
	}
	if err := lru.Put("5", 5); err != nil {
		panic(err)
	}
	if err := lru.Put("5", 6); err != nil {
		fmt.Println(err)
	}
	if err := lru.Update("5", 6); err != nil {
		fmt.Println(err)
	}
	if err := lru.Update("6", 6); err != nil {
		fmt.Println(err)
	}
	s := []string{"2", "3", "4", "5"}
	for _, k:= range s{
		if value, err := lru.Get(k); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d\n", value)
		}

	}
	lru.String()
}

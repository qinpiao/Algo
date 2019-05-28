package linklist

import (
	"fmt"
	"errors"
)

type node struct {
	next *node
	data interface{}
}

type linkList struct {
	length int
	head *node
}

func NewLinkList(v ...interface{}) *linkList {
	l := &linkList{length:0, head:&node{next: nil}}
	nextNode := l.head
	for _, e := range v {
		nextNode.next = &node{data:e}
		l.length++
		nextNode = nextNode.next
	}
	return l
}

func (l *linkList) String() {
	nextNode := l.head.next
	for nil != nextNode {
		fmt.Printf("%v ", nextNode.data)
		nextNode = nextNode.next
	}
	fmt.Printf("len: %d\n", l.length)
}

func (l *linkList) Append(data interface{}) {
	nextNode := l.head
	for nil != nextNode.next {
		nextNode = nextNode.next
	}
	nextNode.next = &node{data:data}
	l.length++
}

func (l *linkList) Delete(data interface{}) {
	nextNode := l.head
	for nil != nextNode.next {
		if nextNode.next.data == data {
			nextNode.next, nextNode.next.next = nextNode.next.next, nextNode.next
			l.length--
			break
		}
		nextNode = nextNode.next
	}
}

func (l *linkList) Insert(index int, datas ...interface{}) error {
	if index >= l.length || index < 0 {
		return errors.New("index out of range")
	}
	nextNode := l.head
	for i := 0; i < index; i++ {
		nextNode = nextNode.next
	}
	restNode := nextNode.next
	for _, data := range datas {
		nextNode.next = &node{data:data}
		l.length ++
		nextNode = nextNode.next
	}
	nextNode.next = restNode
	return nil
}

func (l *linkList) Get(index int) (interface{}, error) {
	if index >= l.length || index < 0 {
		return nil, errors.New("index out of range")
	}
	nextNode := l.head
	for i := 0; i < index; i++ {
		nextNode = nextNode.next
	}
	return nextNode.next.data, nil
}

func (l *linkList) Reverse() *linkList {
	nl := NewLinkList()
	nl.length = l.length
	head := nl.head
	for NextNode := l.head.next;NextNode != nil; NextNode = NextNode.next{
		newNode := &node{data:NextNode.data, next:head.next}
		head.next = newNode
	}
	return nl
}
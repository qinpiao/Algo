package tree

import (
	"fmt"
)

type node struct {
	data interface{}
	left *node
	right *node
}

type binaryTree struct {
	root *node
}

func NewNode(data interface{}) *node {
	newNode := &node{data:data}
	return newNode
}

func NewBinaryTree(rootv interface{}) *binaryTree {
	newTree := &binaryTree{root:&node{data: rootv}}
	return newTree
}

func preOrderNext(n *node){
	fmt.Println(n.data)
	if n.left != nil {
		preOrderNext(n.left)
	}
	if n.right != nil {
		preOrderNext(n.right)
	}
}

func (bt *binaryTree) PreOrderTraverse(){
	preOrderNext(bt.root)
}

func inOrderNext(n *node){
	if n.left != nil {
		inOrderNext(n.left)
	}
	fmt.Println(n.data)
	if n.right != nil {
		inOrderNext(n.right)
	}
}

func (bt *binaryTree) InOrderTraverse(){
	inOrderNext(bt.root)
}

func postOrderNext(n *node){
	if n.left != nil {
		postOrderNext(n.left)
	}
	if n.right != nil {
		postOrderNext(n.right)
	}
	fmt.Println(n.data)
}

func (bt *binaryTree) PostOrderTraverse(){
	postOrderNext(bt.root)
}

func hierarchicalNext(queue []*node) {
	newQueue := []*node{}
	for _, n := range queue {
		fmt.Printf("%d ", n.data)
		if nil != n.left {
			newQueue = append(newQueue, n.left)
		}
		if nil != n.right {
			newQueue = append(newQueue, n.right)
		}
	}
	fmt.Println()
	if len(newQueue) > 0 {
		hierarchicalNext(newQueue)
	}
}

func (bt *binaryTree) HierarchicalTraverse(){
	queue := []*node{bt.root}
	hierarchicalNext(queue)
}

func zNext(queue []*node, reverse bool){
	newQueue := []*node{}
	if len(queue) > 0 {
		if reverse {
			for _, n := range queue {
				fmt.Printf("%d ", n.data)
				if nil != n.left {
					newQueue = append(newQueue, n.left)
				}
				if nil != n.right {
					newQueue = append(newQueue, n.right)
				}
			}
		} else {
			for i := len(queue)-1; i>=0; i-- {
				n := queue[i]
				fmt.Printf("%d ", n.data)
				if nil != n.left {
					newQueue = append(newQueue, n.left)
				}
				if nil != n.right {
					newQueue = append(newQueue, n.right)
				}

			}
		}
	}
	fmt.Println()
	reverse = !reverse
	if len(newQueue) > 0 {
		zNext(newQueue, reverse)
	}
}

func (bt *binaryTree) ZTraverse() {
	queue := []*node{bt.root}
	zNext(queue, true)
}
package tree

import (
	"testing"
	"fmt"
)

func IntCompare (v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)
	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}
	return 0
}


func TestBinarySearchTree(t *testing.T) {
	s := []int{4,5,2,1,6,7,8}
	bst := NewBST(3, IntCompare)
	for _, e := range s {
		if err := bst.Add(e); err != nil {
			panic(err)
		}
	}
	bst.InOrderTraverse()
	fmt.Println()
	bst.PreOrderTraverse()
	fmt.Println()
	bst.PostOrderTraverse()
	fmt.Println()
	fmt.Println(bst.Min())
	fmt.Println(bst.Max())
	bst.HierarchicalTraverse()
	fmt.Println()
	bst.ZTraverse()
	fmt.Println()
	n := bst.Find(2)
	if nil != n {
		fmt.Println(n.data)
	}

	bst.Delete(2)
	bst.InOrderTraverse()
	fmt.Println()
	bst.HierarchicalTraverse()
	fmt.Println()
	bst.Mirror()


}


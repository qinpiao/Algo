package tree

import (
	"fmt"
	"errors"
)

type CompareFunc func(v, nodeV interface{}) int

type BST struct{
	*binaryTree
	Compare CompareFunc
}

func NewBST(rootv interface{}, comparefunc CompareFunc) *BST {
	bt := NewBinaryTree(rootv)
	bst := &BST{binaryTree:bt, Compare: comparefunc}
	return bst
}

func (bt *BST) Add(v interface{}) error {
	n := bt.root
	for n != nil {
		switch bt.Compare(v, n.data) {
		case 0 : {
			return errors.New(fmt.Sprintf("%v Already exists", v))
		}
		case 1 : {
			if nil == n.right {
				n.right = NewNode(v)
				n = nil
			} else {
				n = n.right
			}

		}
		case -1 : {
			if nil == n.left {
				n.left = NewNode(v)
				n = nil
			} else {
				n = n.left
			}

		}
		}
	}
	return nil
}

func (bt *BST) Find(v interface{}) *node {
	n := bt.root
	for nil != n {
		switch bt.Compare(v, n.data) {
		case 0: {
			return n
		}
		case 1: {
			n = n.right
		}
		case -1: {
			n = n.left
		}
		}
	}
	return nil
}

func (bt *BST) Delete(v interface{}) error {
	var prevNode, targetNode *node = nil, nil
	var side string
	// 先找到要删除的node的位置
	n := bt.root
	for nil != n {
		switch bt.Compare(v, n.data) {
		case 0: {
			targetNode = n
			n = nil
		}
		case 1: {
			prevNode = n
			side = "r"
			n = n.right
		}
		case -1: {
			prevNode = n
			side = "l"
			n = n.left
		}
		}
	}
	if nil == targetNode {
		return errors.New(fmt.Sprintf("%v Not Exists", v))
	}

	// 整棵树只有根节点的情况
	if nil == prevNode && targetNode != nil {
		targetNode.data = nil
		return nil
	}
	// 待删除节点为叶子节点
	if nil == targetNode.left && nil == targetNode.right {
		switch side {
		case "l": prevNode.left = nil
		case "r": prevNode.right = nil
		}
	} else if nil != targetNode.left && nil != targetNode.right {
		mn := targetNode.right
		prevmn := targetNode
		for nil != mn {
			mn = mn.left
		}
		switch side {
		case "l": {
			prevNode.left = mn
			prevmn.left = nil
		}
		case "r": {
			prevNode.right = mn
			prevmn.left = mn
		}
		}
	} else if nil == targetNode.left || nil == targetNode.right {
		switch side {
		case "l": {
			if nil == targetNode.left {
				prevNode.left = targetNode.right
			} else {
				prevNode.left = targetNode.left
			}
		}
		case "r": {
			if nil == targetNode.left {
				prevNode.right = targetNode.right
			} else {
				prevNode.right = targetNode.left
			}
		}
		}

	}
	return nil
}

func (bt *BST) Min() interface{}{
	n := bt.root
	var target *node
	for n != nil {
		target = n
		n = n.left
	}
	return target.data
}

func (bt *BST) Max() interface{}{
	n := bt.root
	var target *node
	for n != nil {
		target = n
		n = n.right
	}
	return target.data
}
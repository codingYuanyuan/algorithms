package BSTsort

import (
	"fmt"
)

type Node struct {
	Key   int
	N     int
	Left  *Node
	Right *Node
	Value int
}

func getsize(root *Node) int {
	if root == nil {
		return 0
	}
	return root.N
}

func Getsize(root *Node) int {
	return root.N
}

func Put(inpN int, root *Node, value int) *Node {
	if root == nil {
		return &Node{inpN, 1, nil, nil, value}
	}
	if inpN < root.Key {
		//fmt.Println("go left", root.Value, "current value", value)
		root.Left = Put(inpN, root.Left, value)
	} else if inpN > root.Key {
		//fmt.Println("go right", root.Value, "current value", value)
		root.Right = Put(inpN, root.Right, value)
	} else {
		root.Value = value
	}
	root.N = getsize(root.Left) + getsize(root.Right) + 1

	return root
}

func Print(root *Node) {
	if root != nil {

		Print(root.Left)
		fmt.Println(root.Value)
		Print(root.Right)
	}
}

func Print1(root *Node) {
	if root != nil {
		fmt.Println(root.Value)
		Print1(root.Left)
		Print1(root.Right)
	}
}
func Print2(root *Node) {
	if root != nil {
		Print2(root.Left)
		Print2(root.Right)
		fmt.Println(root.Value)
	}
}

func Find(key int, root *Node) int {
	if root == nil {
		return -1000
	}
	if key < root.Key {
		return Find(key, root.Left)
	} else if key > root.Key {
		return Find(key, root.Right)
	}
	return root.Value
}

func Min(root *Node) (*Node, bool) {
	if root == nil {
		return &Node{-1000, -10000, nil, nil, -10000}, false
	}
	if root.Left == nil {
		return root, true
	}
	node, ok := Min(root.Left)
	return node, ok
}

func Max(root *Node) (*Node, bool) {
	if root == nil {
		return &Node{-1000, -10000, nil, nil, -10000}, false
	}
	if root.Right == nil {
		return root, true
	}
	node, ok := Max(root.Right)
	return node, ok
}

func Floor(key int, root *Node) (int, bool) {
	if root == nil {
		return -10000000, false
	}
	if key == root.Key {
		return key, true
	}
	if key < root.Key {
		return Floor(key, root.Left)
	}
	x, ok := Floor(key, root.Right)

	if ok != true {
		return key, true
	}

	return x, true
}

func Rank(key int, root *Node) int {
	if root == nil {
		return 0
	}
	if key < root.Key {
		return Rank(key, root.Left)
	} else if key > root.Key {
		return 1 + getsize(root.Left) + Rank(key, root.Right)
	}
	return getsize(root.Left)

}

func deletemin(node *Node) *Node {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deletemin(node.Left)
	node.N = getsize(node.Left) + getsize(node.Right) + 1
	return node
}

func Delete(node *Node, key int) (*Node, bool) {
	if node == nil {
		return &Node{-1000, -10000, nil, nil, -10000}, false
	}
	if key < node.Key {
		node, _ = Delete(node.Left, key)
	} else if key > node.Key {
		node, _ = Delete(node.Right, key)
	}
	if node.Left == nil {
		return node.Right, true
	}
	if node.Right == nil {
		return node.Left, true
	}
	temp := node
	node, ok := Min(temp.Right)
	node.Right = deletemin(temp.Right)
	node.Left = temp.Left

	node.N = getsize(node.Left) + getsize(node.Right) + 1
	return node, ok

}

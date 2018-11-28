package main

import (
	"fmt"
	"mywork/algorithms/searching/BSTsort"
)

func main() {
	var root *BSTsort.Node
	root = BSTsort.Put(10, root, 3)
	root = BSTsort.Put(23, root, 6)
	root = BSTsort.Put(13, root, 4)
	root = BSTsort.Put(4, root, 5)
	root = BSTsort.Put(5, root, 7)
	root = BSTsort.Put(11, root, 10)
	root = BSTsort.Put(7, root, 11)
	root = BSTsort.Put(12, root, 2)
	root = BSTsort.Put(9, root, 8)
	root = BSTsort.Put(21, root, 9)

	BSTsort.Print(root)
	fmt.Println("---------")
	BSTsort.Print1(root)
	fmt.Println("---------")

	BSTsort.Print2(root)

	fmt.Println("size: ", BSTsort.Getsize(root))
	fmt.Println(BSTsort.Find(8, root))
	fmt.Println(BSTsort.Find(9, root))
	fmt.Println(BSTsort.Find(7, root))
	fmt.Println(BSTsort.Find(100, root))
	min, ok := BSTsort.Min(root)
	if ok {
		fmt.Println("min: ", min.Key, min.Value)
	}
	max, ok := BSTsort.Max(root)
	if ok {
		fmt.Println("max: ", max.Key, max.Value)
	}
	a, _ := BSTsort.Floor(6, root)
	fmt.Println("floor: ", a)
	b, _ := BSTsort.Floor(1000, root)
	fmt.Println("floor: ", b)
	fmt.Println("Rank", 8, BSTsort.Rank(8, root))
	k, ok := BSTsort.Delete(root, 10)
	if ok {
		fmt.Println("delete")
		BSTsort.Print(k)
	}
}

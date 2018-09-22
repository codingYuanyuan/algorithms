package main

import (
	"fmt"
	"mywork/algorithms/sorting/BSTsort"
)

func main() {
	a := []int{10, 4, 12, 345, 32, 2, 45, 3}
	//c := mergesort.Merges(a)
	//c := insertsort.Insertionsort(a)
	//c := heapsort.Heap(a)
	//c = heapsort.Increasekey(a, 2, 100)
	//c = heapsort.ExtractMax(a)
	//c := heapsort.HeapSort(a)
	var tree *BSTsort.Root
	var tree1 *BSTsort.Root

	for len(a) > 0 {
		tree = BSTsort.Insert(tree, a[0])
		tree1 = BSTsort.Insert2(tree1, a[0])
		a = a[1:]
	}
	BSTsort.PrintTree(tree)
	fmt.Println("next")
	BSTsort.PrintTree(tree1)
	fmt.Println("In order walk")

	BSTsort.Inorderwalk(tree)
}

package main

import (
	"fmt"
	"mywork/algorithms/sorting/heapsort"
)

func main() {
	a := []int{3951, 10, 4, 12, 345, 32, 2, 45}
	//c := mergesort.Merges(a)
	//c := quicksort.Quicksort(a)
	c := heapsort.HeapSort(a)
	//c = heapsort.AddKey(a, 200)
	//c := mergesort.BUmerge(a)
	//c := insertsort.Insertionsort(a)
	//c := heapsort.Heap(a)
	//c = heapsort.Increasekey(a, 2, 100)
	//c = heapsort.ExtractMax(a)
	//c := heapsort.HeapSort(a)
	//var tree *BSTsort.Root
	//var tree1 *BSTsort.Root
	//c := selectionsort.Selectionsort(a)
	//c := shellsort.ShellSort(a)
	fmt.Println(c)

	/***for len(a) > 0 {
		tree = BSTsort.Insert(tree, a[0])
		tree1 = BSTsort.Insert2(tree1, a[0])
		a = a[1:]
	}

	BSTsort.PrintTree(tree)
	fmt.Println("next")
	BSTsort.PrintTree(tree1)
	fmt.Println("In order walk")

	BSTsort.Inorderwalk(tree)
	fmt.Println("Preorder walk")
	BSTsort.Preorderwalk(tree)
	fmt.Println("Post order walk")
	BSTsort.Postorderwalk(tree)
	***/
}

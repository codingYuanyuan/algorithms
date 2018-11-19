package heapsort

import "fmt"

func HeapSort(a []int) []int {
	start := len(a)/2 - 1

	//build up the heap
	for start >= 0 {
		//fmt.Println(a, start)
		a = maxHeapify(a, len(a), start)
		start = start - 1
	}
	//sorting
	//fmt.Println("start sorting")
	ind := len(a) - 1
	for ind >= 0 {
		temp := a[0]
		a[0] = a[ind]
		a[ind] = temp
		a = maxHeapify(a, ind, 0)
		ind = ind - 1
		//fmt.Println("Next ind")
	}
	return a
}

func maxHeapify(a []int, b int, c int) []int {
	Left := left(b, c)
	Right := right(b, c)
	lar := c
	if Left != -1 && a[Left] > a[lar] {
		lar = Left
	}
	if Right != -1 && a[Right] > a[lar] {
		lar = Right
	}
	//fmt.Println(a, c, lar, "inner")
	if lar != c {
		a = exch(a, lar, c)
		a = maxHeapify(a, b, lar)
	}

	return a
}

func exch(a []int, b int, c int) []int {
	tem := a[b]
	a[b] = a[c]
	a[c] = tem
	return a
}
func left(n int, p int) int {
	if 2*p+1 >= n {
		return -1
	}
	return 2*p + 1
}

func right(n int, p int) int {
	if 2*p+2 >= n {
		return -1
	}
	return 2*p + 2
}

func DeleteKey(a []int, keyind int) []int {
	temp := a[0]
	a[0] = a[len(a)-1]
	a[len(a)-1] = temp
	kl := len(a) - 1
	a = a[:len(a)-1]
	a = sink(a, kl, 0)
	return a

}

func AddKey(a []int, key int) []int {
	a = append(a, key)
	fmt.Println(a)
	a = swim(a, len(a)-1)
	return a
}

func swim(a []int, k int) []int {
	for k > 1 && a[k/2] > a[k] {
		a = exch(a, k/2, k)
		k = k / 2
	}
	return a
}

func sink(a []int, k int, c int) []int {
	for 2*c <= k {
		j := 2 * c
		fmt.Println(j)
	}
	return a
}

/***
func Heap(a []int) []int {
	start := len(a) / 2
	for start >= 0 {
		a = MaxHeapify(a, start)
		start = start - 1
	}
	return a
}

func MaxHeapify(a []int, b int) []int {
	lef := left(a, b)
	rig := right(a, b)

	largest := b
	if rig != -1 && a[rig] > a[largest] {
		largest = rig
	}
	if lef != -1 && a[lef] > a[largest] {
		largest = lef
	}
	//fmt.Println("what", largest == b)

	if largest != b {
		a = exchange(a, largest, b)
		a = MaxHeapify(a, largest)
	}
	return a

}

func exchange(a []int, b int, c int) []int {
	temp := a[b]
	a[b] = a[c]
	a[c] = temp
	return a
}

func left(a []int, i int) int {
	le := -1
	if i == 0 {
		if len(a) > 1 {
			le = 1
		}
	} else {
		if len(a) > 2*i {
			le = 2 * i
		}
	}
	return le
}

func right(a []int, i int) int {
	ri := -1
	if i == 0 {
		if len(a) > 2 {
			ri = 2
		}
	} else {
		if len(a) > 2*i+1 {
			ri = 2*i + 1
		}
	}
	return ri
}

func Increasekey(a []int, i int, value int) []int {
	if value < a[i] {
		return a
	}
	if i >= len(a) {
		return a
	}
	a[i] = value
	return Heap(a)
}

func ExtractMax(a []int) []int {
	a = a[1:]
	a = Heap(a)
	return a
}

func HeapSort(a []int) []int {
	a = Heap(a)
	start := len(a) - 1
	fina := len(a) - 1
	for start > 0 {
		a = exchange(a, 0, start)
		b := a[:fina]
		c := a[fina:]
		b = MaxHeapify(b, 0)
		a = appendslice(b, c)
		fina = fina - 1
		start = start - 1
	}
	return a

}

func appendslice(a []int, b []int) []int {
	for len(b) != 0 {
		a = append(a, b[0])
		b = b[1:]
	}
	return a
}
***/

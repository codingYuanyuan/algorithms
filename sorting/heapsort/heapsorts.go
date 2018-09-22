package heapsort

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

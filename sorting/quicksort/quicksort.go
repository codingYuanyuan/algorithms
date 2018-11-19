package quicksort

func Quicksort(a []int) []int {
	a = tsort(a, 0, len(a)-1)
	return a
}
func sort(a []int, lo int, hi int) []int {
	if hi <= lo {
		return a
	}
	j, a := partition(a, lo, hi)
	//fmt.Println(j, a, "start", lo, hi)
	a = sort(a, lo, j-1)
	//fmt.Println(lo, j-1, a, " right")
	a = sort(a, j+1, hi)
	//fmt.Println(j+1, hi, a, "left")
	return a
}

func tsort(a []int, lo int, hi int) []int {
	if hi <= lo {
		return a
	}
	lt := lo
	gt := hi
	i := lo + 1
	v := a[lo]
	for i <= gt {
		if a[i] < v {
			a = exch(a, i, lt)
			i = i + 1
			lt = lt + 1
		} else if a[i] > v {
			a = exch(a, i, gt)
			gt = gt - 1
		} else {
			i = i + 1
		}
	}
	a = tsort(a, lo, lt-1)
	a = tsort(a, gt+1, hi)
	return a
}

func partition(a []int, lo int, hi int) (int, []int) {
	i := lo + 1
	j := hi
	v := a[lo]
	for i < j {
		for a[i] < v {
			i = i + 1
			if i == hi {
				break
			}
		}
		for a[j] > v {
			j = j - 1

			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a = exch(a, i, j)

	}
	a = exch(a, lo, j)

	return j, a
}

func exch(a []int, i int, j int) []int {
	temp := a[j]
	a[j] = a[i]
	a[i] = temp
	return a
}

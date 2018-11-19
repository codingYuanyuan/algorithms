package mergesort

func Merges(a []int) []int {
	a = sort(a, 0, len(a)-1)
	return a
}

func sort(a []int, lo int, hi int) []int {
	if hi <= lo {
		return a
	}
	mid := lo + (hi-lo)/2
	a = sort(a, lo, mid)
	a = sort(a, mid+1, hi)
	a = merge(a, lo, mid, hi)
	return a
}

func BUmerge(a []int) []int {
	sz := 1
	for sz < len(a) {
		ind := 0
		for ind < len(a)-sz {
			min := ind + 2*sz - 1
			if min > len(a)-1 {
				min = len(a) - 1
			}
			//fmt.Println(a, ind, ind+sz-1, min)
			a = merge(a, ind, ind+sz-1, min)
			ind += 2 * sz
		}
		sz = sz * 2
	}
	return a
}

func merge(a []int, lo int, mid int, hi int) []int {
	N := len(a)
	var aux []int
	aux = make([]int, N)
	t := 0
	for t < len(a) {
		aux[t] = a[t]
		t += 1
	}

	k := lo
	i := lo
	j := mid + 1

	for k <= hi {
		if i > mid {
			aux[k] = a[j]
			j += 1
		} else if j > hi {
			aux[k] = a[i]
			i += 1

		} else if a[i] < a[j] {
			aux[k] = a[i]
			i += 1
		} else {
			aux[k] = a[j]
			j += 1
		}
		k += 1
	}
	return aux
}

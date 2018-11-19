package selectionsort

func Selectionsort(a []int) []int {
	N := len(a)
	i := 0
	for i < N {
		min := i
		j := i + 1
		for j < N {
			if a[min] > a[j] {
				min = j
			}
			a = exch(a, i, min)
			j += 1
		}
		i += 1
	}
	return a
}

func exch(a []int, b int, c int) []int {
	temp := a[b]
	a[b] = a[c]
	a[c] = temp
	return a
}

package shellsort

func ShellSort(a []int) []int {
	N := len(a)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		i := h
		for i < N {
			st := i
			for st >= h {
				if a[st] < a[st-h] {
					a = exch(a, st, st-h)

				}
				st -= h

			}
			//fmt.Println(h)
			//fmt.Println(a)
			i += 1
		}

		h = h / 3
	}
	return a
}

func exch(a []int, b int, c int) []int {
	temp := a[b]
	a[b] = a[c]
	a[c] = temp
	return a
}

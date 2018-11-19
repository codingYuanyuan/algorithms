package insertsort

func switchnum(a []int, ind1 int, ind2 int) []int {
	temp := a[ind1]
	a[ind1] = a[ind2]
	a[ind2] = temp
	return a
}

func Insertionsort(a []int) []int {
	ind := 1

	for ind < len(a) {
		poin := ind
		for poin > 0 {
			if a[poin] < a[poin-1] {
				a = exch(a, poin, poin-1)
			}
			poin -= 1
			//fmt.Println(ind)
			//fmt.Println(poin)
			//fmt.Println(a)
		}
		ind += 1
	}

	return a
}

func exch(a []int, b int, c int) []int {
	temp := a[b]
	a[b] = a[c]
	a[c] = temp
	return a
}

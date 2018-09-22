package selectionsort

func switchnum(a []int, ind1 int, ind2 int) []int {
	temp := a[ind1]
	a[ind1] = a[ind2]
	a[ind2] = temp
	return a
}

func Insertionsort(a []int) []int {
	ind := 0

	for ind <= len(a)-1 {
		poin := ind
		for poin <= len(a)-1 {
			if a[poin] < a[ind] {
				a = switchnum(a, ind, poin)
			}
			poin = poin + 1
		}
		ind = ind + 1
	}
	return a
}

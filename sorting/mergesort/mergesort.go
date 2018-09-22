package mergesort

func Merges(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	l := len(a) / 2
	b := Merges(a[:l])
	c := Merges(a[l:])

	d := merge(b, c)

	return d
}

func merge(a []int, b []int) []int {
	c := []int{}

	for len(a) > 0 && len(b) > 0 {
		if a[0] >= b[0] {
			c = append(c, b[0])
			b = b[1:]
		} else {
			c = append(c, a[0])
			a = a[1:]
		}
	}

	if len(a) != 0 {
		for len(a) > 0 {
			c = append(c, a[0])
			a = a[1:]
		}
	}

	if len(b) != 0 {
		for len(b) > 0 {
			c = append(c, b[0])
			b = b[1:]
		}
	}

	return c
}

package BSTsort

import (
	"fmt"
)

type Root struct {
	Value int
	Left  *Root
	Right *Root
}

func Insert2(r *Root, a int) *Root {
	if r == nil {
		return &Root{a, nil, nil}
	}
	if a < r.Value {
		r.Left = Insert2(r.Left, a)
	} else {
		r.Right = Insert2(r.Right, a)
	}
	return r
}

func Insert(r *Root, a int) *Root {
	if r == nil {
		return &Root{a, nil, nil}
	}

	count := 0
	base := r

	for r.Left != nil && r.Right != nil {
		if count == 0 {
			base = r
		}
		if a < r.Value {
			r = r.Left
		} else {
			r = r.Right
		}
		count += 1
	}

	for r.Left != nil {
		if count == 0 {
			base = r
		}
		if a < r.Value {
			r = r.Left
		} else {
			if r.Right != nil {
				r = r.Right
			} else {
				r.Right = &Root{a, nil, nil}
				return base
			}
		}
		count += 1
	}

	for r.Right != nil {
		if count == 0 {
			base = r
		}
		if a > r.Value {
			r = r.Right
		} else {
			if r.Left != nil {
				r = r.Left
			} else {
				r.Left = &Root{a, nil, nil}
				return base
			}
		}
		count += 1
	}

	if a < r.Value {
		r.Left = &Root{a, nil, nil}
	} else {
		r.Right = &Root{a, nil, nil}
	}
	return base
}

func PrintTree(r *Root) {
	if r != nil {
		fmt.Println(r.Value)
		PrintTree(r.Left)
		PrintTree(r.Right)
	}
}

func Inorderwalk(r *Root) {
	if r != nil {
		Inorderwalk(r.Left)
		fmt.Println(r.Value)
		Inorderwalk(r.Right)
	}
}

func Preorderwalk(r *Root) {
	if r != nil {
		fmt.Println(r.Value)
		Preorderwalk(r.Left)
		Preorderwalk(r.Right)
	}
}

func Postorderwalk(r *Root) {
	if r != nil {

		Preorderwalk(r.Left)
		Preorderwalk(r.Right)
		fmt.Println(r.Value)
	}
}

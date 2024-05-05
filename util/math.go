package util

import "golang.org/x/exp/constraints"

func Min[T Number](x, y T) T {
	if x < y {
		return x
	}

	return y
}

type Number interface {
	constraints.Integer | constraints.Float
}

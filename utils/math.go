package utils

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

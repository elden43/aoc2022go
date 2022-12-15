package utils

import "strconv"

func Str2int(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

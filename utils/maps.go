package utils

import (
	"fmt"
	"sort"
)

func DebugSurface(s map[int]map[int]rune) {
	var yc []int
	for y, _ := range s {
		yc = append(yc, y)
	}

	sort.Ints(yc)
	for _, vy := range yc {
		var xc []int
		for x, _ := range s[vy] {
			xc = append(xc, x)
		}
		sort.Ints(xc)

		for _, vx := range xc {
			fmt.Println(vy, vx, ":", string(s[vy][vx]))
		}
	}
}

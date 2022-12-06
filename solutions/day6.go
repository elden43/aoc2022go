package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
)

func Day6part1() {
	for _, line := range reader.ReadLines("puzzles/day6/input.txt") {
		for i := 3; i < len(line); i++ {
			if i < 3 {
				continue
			}
			s := line[i-3 : i+1]
			if unique(s, 4) {
				fmt.Println(i + 1)
				break
			}
		}
	}
}

func Day6part2() {
	for _, line := range reader.ReadLines("puzzles/day6/input.txt") {
		for i := 13; i < len(line); i++ {
			if i < 13 {
				continue
			}
			s := line[i-13 : i+1]
			if unique(s, 14) {
				fmt.Println(i + 1)
				break
			}
		}
	}
}

func unique(s string, count int) bool {
	m := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		m[string(s[i])] = true
	}
	if len(m) == count {
		return true
	}

	return false
}

package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"regexp"
	"strconv"
)

type Range struct {
	min int
	max int
}

func Day4part1() {
	res := 0
	for _, v := range reader.ReadLines("puzzles/day4/input.txt") {
		if fullContains(parseRange(v)) {
			res++
		}
	}
	fmt.Println(res)
}

func Day4part2() {
	res := 0
	for _, v := range reader.ReadLines("puzzles/day4/input.txt") {
		if overlaps(parseRange(v)) {
			res++
		}
	}
	fmt.Println(res)
}

func fullContains(a, b Range) bool {
	if (a.min <= b.min && a.max >= b.max) || (b.min <= a.min && b.max >= a.max) {
		return true
	} else {
		return false
	}
}

func overlaps(a, b Range) bool {
	if a.max >= b.min && a.min <= b.max {
		return true
	}

	return false
}

func parseRange(s string) (a, b Range) {
	re := regexp.MustCompile(`(?m)^([[:digit:]]*)-([[:digit:]]*),([[:digit:]]*)-([[:digit:]]*)$`)
	match := re.FindStringSubmatch(s)
	v1, _ := strconv.Atoi(match[1])
	v2, _ := strconv.Atoi(match[2])
	v3, _ := strconv.Atoi(match[3])
	v4, _ := strconv.Atoi(match[4])

	return Range{v1, v2}, Range{v3, v4}
}

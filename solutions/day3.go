package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
)

func Day3part1() {
	priorityScore := 0
	for _, v := range reader.ReadLines("puzzles/day3/input.txt") {
		comA := v[:len(v)/2]
		comB := v[len(v)/2:]
		item, _ := hashIntersectOne(comA, comB)
		priorityScore += priority(item)
	}

	fmt.Println(priorityScore)
}

func Day3part2() {
	priorityScore := 0
	i := 0
	var lines [3]string

	for _, v := range reader.ReadLines("puzzles/day3/input.txt") {
		lines[i] = v
		i++
		if i == 3 {
			priorityScore += priority(mostOccurrencesOneByLine(lines))
			i = 0
		}
	}

	fmt.Println(priorityScore)
}

func hashIntersectOne(s1, s2 string) (rune, bool) {
	def := rune(0)
	r1 := []rune(s1)
	r2 := []rune(s2)
	hash := make(map[rune]int)

	for i := 0; i < len(r1); i++ {
		item, ok := hash[r1[i]]
		if !ok {
			hash[r1[i]] = 1
		} else {
			hash[r1[i]] = item + 1
		}
	}

	for i := 0; i < len(r2); i++ {
		_, ok := hash[r2[i]]
		if ok {
			return r2[i], true
		}
	}

	return def, false
}

func mostOccurrencesOneByLine(lines [3]string) rune {
	hash := make(map[rune]int)
	for _, v := range lines {
		hashForLine := make(map[rune]bool)
		for i := 0; i < len(v); i++ {
			_, okForLine := hashForLine[rune(v[i])]
			if okForLine {
				continue
			} else {
				hashForLine[rune(v[i])] = true
			}
			item, ok := hash[rune(v[i])]
			if !ok {
				hash[rune(v[i])] = 1
			} else {
				r := rune(v[i])
				hash[r] = item + 1
			}
		}
	}

	maxValue := 0
	maxIndex := '0'
	for i, v := range hash {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}

	return maxIndex
}

func priority(r rune) int {
	if r >= 97 {
		return int(r - 96)
	} else {
		return int(r - 38)
	}
}

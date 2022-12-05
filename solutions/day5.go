package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"regexp"
	"strconv"
	"strings"
)

func Day5part1() {
	piles := pilesMap()
	doMoving(piles, true)

	res := ""
	for i := 1; i <= len(piles); i++ {
		res = res + string((*piles[i])[len(*piles[i])-1])
	}

	fmt.Println(res)

}

func Day5part2() {
	piles := pilesMap()
	doMoving(piles, false)

	res := ""
	for i := 1; i <= len(piles); i++ {
		res = res + string((*piles[i])[len(*piles[i])-1])
	}

	fmt.Println(res)
}

func pilesMap() map[int]*string {
	piles := make(map[int]*string)

	for _, v := range reader.ReadLines("puzzles/day5/input.txt") {
		if !strings.Contains(v, "[") {
			break
		}
		line := []rune(v)
		for il := 1; il < len(line); il = il + 4 {
			if line[il] != ' ' {
				index := (il / 4) + 1

				tmp := ""
				val, ok := piles[index]
				if ok {
					tmp = *val
				}
				nval := string(line[il]) + tmp
				piles[index] = &nval
			}
		}
	}

	return piles
}

func doMoving(piles map[int]*string, reverseOrder bool) {
	for _, v := range reader.ReadLines("puzzles/day5/input.txt") {
		if strings.Contains(v, "move") {
			re := regexp.MustCompile(`(?m)^move ([[:digit:]]*) from ([[:digit:]]*) to ([[:digit:]]*)`)
			match := re.FindStringSubmatch(v)
			v1, _ := strconv.Atoi(match[1])
			v2, _ := strconv.Atoi(match[2])
			v3, _ := strconv.Atoi(match[3])
			move(piles, v1, v2, v3, reverseOrder)
		}
	}
}

func move(piles map[int]*string, count int, from int, to int, reverseOrder bool) {
	sourcePileStr := *piles[from]
	nSourcePile := sourcePileStr[0 : len(sourcePileStr)-count]
	piles[from] = &nSourcePile

	moving := sourcePileStr[len(sourcePileStr)-count:]

	destPileStr := *piles[to]
	if reverseOrder {
		destPileStr = destPileStr + reverseString(moving)
	} else {
		destPileStr = destPileStr + moving
	}
	piles[to] = &destPileStr
}

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

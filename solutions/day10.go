package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"strconv"
	"strings"
)

type cycleValues struct {
	before int
	after  int
}

func newCycleValues(before, after int) *cycleValues {
	cv := cycleValues{before, after}
	return &cv
}

func recordCycles(input []string) map[int]*cycleValues {
	cycles := make(map[int]*cycleValues)
	val := 1

	for _, v := range input {
		v = strings.TrimSpace(v)
		if strings.Contains(v, "addx") {
			val = addx(val, v[strings.Index(v, " ")+1:], cycles)
		} else {
			noop(val, cycles)
		}
	}
	return cycles
}

func addx(val int, s string, cycles map[int]*cycleValues) int {
	x, _ := strconv.Atoi(s)

	curIndex := len(cycles)
	cycles[curIndex] = newCycleValues(val, val)
	curIndex++
	cycles[curIndex] = newCycleValues(val, val+x)

	return val + x
}

func noop(val int, cycles map[int]*cycleValues) {
	curIndex := len(cycles)
	cycles[curIndex] = newCycleValues(val, val)
}

func Day10part1() {
	input := reader.GetInputSeparatedBy("puzzles/day10/input.txt", "\n")
	cycles := recordCycles(input)
	res := 0

	for i := 0; i <= len(cycles); i++ {
		if i == 19 || (i-19)%40 == 0 {
			res += (i + 1) * cycles[i].before
		}
	}

	fmt.Println(res)
}

func Day10part2() {
	input := reader.GetInputSeparatedBy("puzzles/day10/input.txt", "\n")
	cycles := recordCycles(input)

	for i := 0; i <= len(cycles)-1; i++ {
		spritePos := i % 40
		if cycles[i].before >= (spritePos-1) && cycles[i].before <= (spritePos+1) {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		if (i+1)%40 == 0 {
			fmt.Println()
		}
	}
}

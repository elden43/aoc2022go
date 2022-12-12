package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"github.com/elden43/aoc2022/utils"
	"strconv"
	"strings"
)

type mapField struct {
	head     bool
	tails    map[int]bool
	tVisited bool
}

func newField(head bool, tVisited bool, tailsCount int) *mapField {
	tails := make(map[int]bool)
	for i := 1; i <= tailsCount; i++ {
		tails[i] = false
	}
	mf := &mapField{head, tails, tVisited}
	return mf
}

func Day9part1() {
	input := reader.GetInputSeparatedBy("puzzles/day9/input.txt", "\n")

	tailsCount := 1
	fields := make(map[int]map[int]*mapField)
	fields[0] = make(map[int]*mapField)
	fields[0][0] = newField(true, true, tailsCount)
	fields[0][0].tails[1] = true

	for _, v := range input {
		dir := v[0:1]
		cnt, _ := strconv.Atoi(strings.TrimSpace(v[2:]))
		moveKnots(dir, cnt, tailsCount, fields)
	}

	visited := 0
	for _, vx := range fields {
		for _, vy := range vx {
			if vy.tVisited == true {
				visited++
			}
		}
	}

	fmt.Println(visited)
}

func Day9part2() {
	input := reader.GetInputSeparatedBy("puzzles/day9/input.txt", "\n")

	tailsCount := 9
	fields := make(map[int]map[int]*mapField)
	fields[0] = make(map[int]*mapField)
	fields[0][0] = newField(true, true, tailsCount)
	for i := 1; i <= tailsCount; i++ {
		fields[0][0].tails[i] = true
	}

	for _, v := range input {
		dir := v[0:1]
		cnt, _ := strconv.Atoi(strings.TrimSpace(v[2:]))
		moveKnots(dir, cnt, tailsCount, fields)
	}

	visited := 0
	for _, vx := range fields {
		for _, vy := range vx {
			if vy.tVisited == true {
				visited++
			}
		}
	}

	fmt.Println(visited)
}

func moveKnots(direction string, count, tailsCount int, fields map[int]map[int]*mapField) {
	curHX, curHY := curHPos(fields)
	for i := 0; i < count; i++ {
		fields[curHX][curHY].head = false
		switch direction {
		case "R":
			curHX++
		case "L":
			curHX--
		case "U":
			curHY++
		case "D":
			curHY--
		}
		verifyMap(curHX, curHY, tailsCount, fields)
		fields[curHX][curHY].head = true

		for ti := 1; ti <= tailsCount; ti++ {
			moveTail(ti, tailsCount, fields)
		}
	}
}

func moveTail(curTail, tailCount int, fields map[int]map[int]*mapField) {
	curHX, curHY := 0, 0
	if curTail == 1 {
		curHX, curHY = curHPos(fields)
	} else {
		curHX, curHY = curTPos(curTail-1, fields)
	}
	curTX, curTY := curTPos(curTail, fields)

	if utils.Abs(curHX-curTX) < 2 && utils.Abs(curHY-curTY) < 2 {
		return
	}
	fields[curTX][curTY].tails[curTail] = false
	if curHX == curTX || curHY == curTY {
		if curHX == curTX {
			curTY = (curHY + curTY) / 2
		} else {
			curTX = (curHX + curTX) / 2
		}
	} else {
		if curHX > curTX {
			curTX++
		} else {
			curTX--
		}
		if curHY > curTY {
			curTY++
		} else {
			curTY--
		}
	}
	verifyMap(curTX, curTY, tailCount, fields)
	fields[curTX][curTY].tails[curTail] = true
	if curTail == tailCount {
		fields[curTX][curTY].tVisited = true
	}
}

func curHPos(fields map[int]map[int]*mapField) (int, int) {
	for ix, vx := range fields {
		for iy, vy := range vx {
			field := vy
			if field.head == true {
				return ix, iy
			}
		}
	}

	return 0, 0
}

func curTPos(tail int, fields map[int]map[int]*mapField) (int, int) {
	for ix, vx := range fields {
		for iy, vy := range vx {
			field := vy
			if field.tails[tail] == true {
				return ix, iy
			}
		}
	}

	return 0, 0
}

func verifyMap(curX, curY, tailsCount int, field map[int]map[int]*mapField) {
	_, okx := field[curX]
	if !okx {
		field[curX] = make(map[int]*mapField)
	}
	_, oky := field[curX][curY]
	if !oky {
		field[curX][curY] = newField(false, false, tailsCount)
	}
}

func showField(field map[int]map[int]*mapField) {
	for ix, vx := range field {
		for iy, vy := range vx {
			xxx := *vy
			for i, _ := range xxx.tails {
				if xxx.tails[i] == true {
					fmt.Println("tail:", ix, iy)
					fmt.Println(xxx.tails)
				}
			}
		}
	}
}

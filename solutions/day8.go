package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
)

func Day8part1() {
	input := reader.GetInputSeparatedBy("puzzles/day8/input.txt", "\n")
	maxX := len(input[0]) - 1
	maxY := len(input)
	trees := make([][]uint8, maxY)
	for y, line := range input {
		trees[y] = make([]uint8, maxX)
		for x := 0; x < maxX; x++ {
			trees[y][x] = line[x]
		}
	}

	visibleCount := maxX*2 + maxY*2 - 4
	for x := 1; x < maxX-1; x++ {
		for y := 1; y < maxY-1; y++ {
			if tallestW(x, y, trees) || tallestE(x, y, trees) || tallestN(x, y, trees) || tallestS(x, y, trees) {
				visibleCount++
			}
		}
	}

	fmt.Println(visibleCount)
}

func Day8part2() {
	input := reader.GetInputSeparatedBy("puzzles/day8/input.txt", "\n")
	maxX := len(input[0]) - 1
	maxY := len(input)
	trees := make([][]uint8, maxY)
	for y, line := range input {
		trees[y] = make([]uint8, maxX)
		for x := 0; x < maxX; x++ {
			trees[y][x] = line[x]
		}
	}

	maxScenicScore := 0
	for x := 1; x < maxX; x++ {
		for y := 1; y < maxY; y++ {
			score := visibleN(x, y, trees) * visibleS(x, y, trees) * visibleW(x, y, trees) * visibleE(x, y, trees)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}

	fmt.Println(maxScenicScore)
}

func visibleW(tx int, ty int, trees [][]uint8) int {
	visible := 0
	if tx == 0 {
		return visible
	}

	for x := tx - 1; x >= 0; x-- {
		visible++
		if trees[ty][x] >= trees[ty][tx] {
			break
		}
	}

	return visible
}

func visibleE(tx int, ty int, trees [][]uint8) int {
	visible := 0
	if tx == len(trees[ty])-1 {
		return visible
	}

	for x := tx + 1; x <= len(trees[ty])-1; x++ {
		visible++
		if trees[ty][x] >= trees[ty][tx] {
			break
		}
	}

	return visible
}

func visibleN(tx int, ty int, trees [][]uint8) int {
	visible := 0
	if ty == 0 {
		return visible
	}

	for y := ty - 1; y >= 0; y-- {
		visible++
		if trees[y][tx] >= trees[ty][tx] {
			break
		}
	}

	return visible
}

func visibleS(tx int, ty int, trees [][]uint8) int {
	visible := 0
	if ty == len(trees)-1 {
		return visible
	}

	for y := ty + 1; y < len(trees); y++ {
		visible++
		if trees[y][tx] >= trees[ty][tx] {
			break
		}
	}

	return visible
}

func tallestW(tx int, ty int, trees [][]uint8) bool {
	tallest := true
	for x := 0; x < tx; x++ {
		if trees[ty][x] >= trees[ty][tx] {
			tallest = false
			break
		}
	}
	return tallest
}

func tallestE(tx int, ty int, trees [][]uint8) bool {
	tallest := true
	for x := len(trees[ty]) - 1; x > tx; x-- {
		if trees[ty][x] >= trees[ty][tx] {
			tallest = false
			break
		}
	}
	return tallest
}

func tallestN(tx int, ty int, trees [][]uint8) bool {
	tallest := true
	for y := 0; y < ty; y++ {
		if trees[y][tx] >= trees[ty][tx] {
			tallest = false
			break
		}
	}
	return tallest
}

func tallestS(tx int, ty int, trees [][]uint8) bool {
	tallest := true
	for y := len(trees) - 1; y > ty; y-- {
		if trees[y][tx] >= trees[ty][tx] {
			tallest = false
			break
		}
	}
	return tallest
}

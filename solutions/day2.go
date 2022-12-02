package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
)

type option int
type result int

const (
	rock option = iota
	paper
	scissors
)

const (
	lose result = iota
	draw
	win
)

var results = map[option]map[option]result{
	rock: map[option]result{
		rock:     draw,
		paper:    lose,
		scissors: win,
	},
	paper: map[option]result{
		rock:     win,
		paper:    draw,
		scissors: lose,
	},
	scissors: map[option]result{
		rock:     lose,
		paper:    win,
		scissors: draw,
	},
}

var scoreForResult = map[result]int {
	lose: 0,
	draw: 3,
	win:  6,
}

var scoreForOption = map[option]int {
	rock:     1,
	paper:    2,
	scissors: 3,
}

var charToOption = map[rune]option{
	'A': rock,
	'B': paper,
	'C': scissors,
	'X': rock,
	'Y': paper,
	'Z': scissors,
}

var optionToResult = map[option]map[result]option{
	rock: map[result]option {
		lose: scissors,
		draw: rock,
		win: paper,
	},
	paper: map[result]option {
		lose: rock,
		draw: paper,
		win: scissors,
	},
	scissors: map[result]option {
		lose: paper,
		draw: scissors,
		win: rock,
	},
}

var charToResult = map[rune]result {
	'X': lose,
	'Y': draw,
	'Z': win,
}

func Day2part1() {
	score := 0
	for _, v := range reader.ReadLines("puzzles/day2/input.txt") {
		runeS := []rune(v)
		score += scoreForResult[results[charToOption[runeS[2]]][charToOption[runeS[0]]]] + scoreForOption[charToOption[runeS[2]]]
	}

	fmt.Println(score)
}

func Day2part2() {
	score := 0
	for _, v := range reader.ReadLines("puzzles/day2/input.txt") {
		runeS := []rune(v)
		score += scoreForResult[results[optionToResult[charToOption[runeS[0]]][charToResult[runeS[2]]]][charToOption[runeS[0]]]] + scoreForOption[optionToResult[charToOption[runeS[0]]][charToResult[runeS[2]]]]
	}

	fmt.Println(score)
}
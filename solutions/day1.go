package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"github.com/elden43/aoc2022/utils"
	"math"
	"strconv"
)

func Day1part1() {
	var caloriesByElf []int
	calories := 0
	for _, v := range reader.ReadLines("puzzles/day1/input.txt") {
		if v == "" {
			caloriesByElf = append(caloriesByElf, calories)
			calories = 0
		} else {
			intVal, _ := strconv.Atoi(v)
			calories += intVal
		}
	}
	caloriesByElf = append(caloriesByElf, calories)

	maxCalories := 0
	for _, v := range caloriesByElf {
		maxCalories = utils.Max(maxCalories, v)
	}

	fmt.Println(maxCalories)
}

func Day1part2() {
	var caloriesByElf []int
	calories := 0
	for _, v := range reader.ReadLines("puzzles/day1/input.txt") {
		if v == "" {
			caloriesByElf = append(caloriesByElf, calories)
			calories = 0
		} else {
			intVal, _ := strconv.Atoi(v)
			calories += intVal
		}
	}
	caloriesByElf = append(caloriesByElf, calories)

	maxCalories := map[int]int{
		0: 0,
		1: 0,
		2: 0,
	}

	for _, v := range caloriesByElf {
		if v > maxCalories[getIndexForLowestValue(maxCalories)] {
			maxCalories[getIndexForLowestValue(maxCalories)] = v
		}
	}

	sum3 := 0
	for _, v := range maxCalories {
		sum3 += v
	}

	fmt.Println(sum3)
}

func getIndexForLowestValue(caloriesMap map[int]int) int {
	lowestIndex := 0
	lowestValue := math.MaxInt
	for i, v := range caloriesMap {
		if v < lowestValue {
			lowestIndex = i
			lowestValue = v
		}
	}

	return lowestIndex
}

package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day11part1() {
	input := reader.GetInputSeparatedBy("puzzles/day11/input.txt", "\r\n\r\n")

	monkeys := make(map[int]*monkey)
	divider := 1
	for _, monkeyParagraph := range input {
		lines := strings.Split(monkeyParagraph, "\r\n")
		monkey := newMonkey(lines)
		monkeys[monkey.id] = monkey
		divider *= monkey.testOperand
	}

	for i := 0; i < 20; i++ {
		round(monkeys, true, divider)
	}

	var inspections []int
	for i := 0; i < len(monkeys); i++ {
		inspections = append(inspections, monkeys[i].inspectedCount)
	}
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	fmt.Println(inspections[0] * inspections[1])
}

func Day11part2() {
	input := reader.GetInputSeparatedBy("puzzles/day11/input.txt", "\r\n\r\n")

	monkeys := make(map[int]*monkey)
	divider := 1
	for _, monkeyParagraph := range input {
		lines := strings.Split(monkeyParagraph, "\r\n")
		monkey := newMonkey(lines)
		monkeys[monkey.id] = monkey
		divider *= monkey.testOperand
	}

	for i := 0; i < 10000; i++ {
		round(monkeys, false, divider)
	}

	var inspections []int
	for i := 0; i < len(monkeys); i++ {
		inspections = append(inspections, monkeys[i].inspectedCount)
	}
	sort.Slice(inspections, func(i, j int) bool {
		return inspections[i] > inspections[j]
	})
	fmt.Println(inspections)
	fmt.Println(inspections[0] * inspections[1])
}

func round(monkeys map[int]*monkey, lowerWorryLevel bool, divider int) {
	for i := 0; i < len(monkeys); i++ {
		itemIds := monkeys[i].getItemIds()
		sort.Ints(itemIds)
		for ii := range itemIds {
			monkeys[i].inspectedCount++
			//raise worry level
			switch monkeys[i].operation {
			case "add":
				monkeys[i].opAddition(ii, monkeys[i].operationOperand, divider)
			case "pow":
				monkeys[i].opSquare(ii, divider)
			case "mul":
				monkeys[i].opMultiplication(ii, monkeys[i].operationOperand, divider)
			}
			//lower worry level
			if lowerWorryLevel {
				monkeys[i].items[ii].lowerWorryLevel()
			}
			//test
			if monkeys[i].items[ii].worryLevel%uint64(monkeys[i].testOperand) == uint64(0) {
				monkeys[i].throwItem(ii, monkeys[i].trueTarget, monkeys)
			} else {
				monkeys[i].throwItem(ii, monkeys[i].falseTarget, monkeys)
			}
		}
	}
}

type monkey struct {
	id               int
	items            map[int]*item
	operation        string
	operationOperand int
	testOperand      int
	trueTarget       int
	falseTarget      int
	inspectedCount   int
}

func (m *monkey) opMultiplication(itemId, multiplier, divider int) {
	m.items[itemId].worryLevel = (m.items[itemId].worryLevel * uint64(multiplier)) % uint64(divider)
}

func (m *monkey) opAddition(itemId, addend, divider int) {
	m.items[itemId].worryLevel = (m.items[itemId].worryLevel + uint64(addend)) % uint64(divider)
}

func (m *monkey) opSquare(itemId, divider int) {
	m.items[itemId].worryLevel = (m.items[itemId].worryLevel * m.items[itemId].worryLevel) % uint64(divider)
}

func (m *monkey) throwItem(itemId, targetMonkeyId int, monkeys map[int]*monkey) {
	//add to target monkey
	monkeys[targetMonkeyId].catchItem(*m.items[itemId])

	//remove from current monkey
	delete(m.items, itemId)
}

func (m *monkey) catchItem(i item) {
	item := i
	m.items[m.newItemId()] = &item
}

func (m *monkey) maxItemId() (maxId int) {
	maxId = -1
	for i := range m.items {
		if i > maxId {
			maxId = i
		}
	}

	return
}

func (m *monkey) newItemId() int {
	return m.maxItemId() + 1
}

func (m *monkey) getItemIds() []int {
	var itemsIds []int
	for i := range m.items {
		itemsIds = append(itemsIds, i)
	}

	return itemsIds
}

func newMonkey(lines []string) *monkey {
	items := make(map[int]*item)
	monkey := monkey{items: items}
	for _, line := range lines {
		//id
		if strings.Contains(line, "Monkey ") {
			id, _ := strconv.Atoi(line[strings.LastIndex(line, " ")+1 : len(line)-1])
			monkey.id = id
		}
		//items
		if strings.Contains(line, "Starting items: ") {
			numbers := strings.Split(line[strings.Index(line, ":")+2:], ",")
			for _, number := range numbers {
				n, _ := strconv.Atoi(strings.TrimSpace(number))
				item := newItem(uint64(n))
				monkey.items[len(monkey.items)] = item
			}
		}
		//operation
		if strings.Contains(line, "Operation: ") {
			//+
			if strings.Contains(line, " + ") {
				n, _ := strconv.Atoi(line[strings.Index(line, "+")+2:])
				monkey.operation = "add"
				monkey.operationOperand = n
			} else if strings.Contains(line, "old * old") {
				//pow
				monkey.operation = "pow"
				monkey.operationOperand = 0
			} else {
				//*
				monkey.operation = "mul"
				n, _ := strconv.Atoi(line[strings.Index(line, "*")+2:])
				monkey.operationOperand = n
			}
		}
		//test
		if strings.Contains(line, "Test: divisible by") {
			n, _ := strconv.Atoi(line[strings.LastIndex(line, " ")+1:])
			monkey.testOperand = n
		}
		//test true
		if strings.Contains(line, "If true: throw to monkey") {
			n, _ := strconv.Atoi(line[strings.LastIndex(line, " ")+1:])
			monkey.trueTarget = n
		}
		//test false
		if strings.Contains(line, "If false: throw to monkey") {
			n, _ := strconv.Atoi(line[strings.LastIndex(line, " ")+1:])
			monkey.falseTarget = n
		}
	}

	return &monkey
}

type item struct {
	worryLevel uint64
}

func (i *item) lowerWorryLevel() {
	i.worryLevel = uint64(math.Floor(float64(i.worryLevel) / float64(3)))
}

func newItem(worryLevel uint64) *item {
	return &item{worryLevel}
}

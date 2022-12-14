package day14

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"github.com/elden43/aoc2022/utils"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	rock       rune = '#'
	entryPoint      = '+'
	activeSand      = 'O'
	sand            = 'o'
	air             = ' '
)

type point struct {
	x int
	y int
}

type helperContainer struct {
	entryPoint       point
	activeSand       bool
	activeSandCoords point
	lowestRockPos    int
	infFloor         bool
	infFloorPos      int
}

func Part1() {
	input := reader.GetInputSeparatedBy("puzzles/day14/input.txt", "\r\n")

	con := helperContainer{
		entryPoint:       point{500, 0},
		activeSand:       false,
		activeSandCoords: point{x: 0, y: 0},
		lowestRockPos:    0,
		infFloorPos:      0,
		infFloor:         false,
	}

	surface := make(map[int]map[int]rune)

	insertPoint(con.entryPoint, surface, entryPoint)
	for _, s := range input {
		insertRocks(s, surface)
	}

	con.lowestRockPos = findLowestRockPos(surface)

	dropout := false
	for !dropout {
		spawnSand(&con, surface)
		dropout = tick(surface, &con)
	}

	fmt.Println(countChar(surface, sand))
}

func Part2() {
	input := reader.GetInputSeparatedBy("puzzles/day14/input.txt", "\r\n")

	con := helperContainer{
		entryPoint:       point{500, 0},
		activeSand:       false,
		activeSandCoords: point{x: 0, y: 0},
		infFloor:         true,
		lowestRockPos:    0,
	}

	surface := make(map[int]map[int]rune)

	for _, s := range input {
		insertRocks(s, surface)
	}

	//infinite floor
	con.infFloorPos = findLowestRockPos(surface) + 2

	blocked := false
	i := 0
	for !blocked {
		i++
		spawnSand(&con, surface)
		_ = tick(surface, &con)
		_, blocked = discover(con.entryPoint, surface)
	}
	fmt.Println(countChar(surface, sand))

}

func countChar(s map[int]map[int]rune, char rune) int {
	count := 0
	for y, yv := range s {
		for x, _ := range yv {
			if s[y][x] == char {
				count++
			}
		}
	}

	return count
}

func tick(s map[int]map[int]rune, con *helperContainer) (dropout bool) {
	if !con.activeSand {
		panic("nononono :-(")
	}

	moved := false
	//shortcut - first drop sand to lowest Y possible
	isUnoccupied := true
	newY := con.activeSandCoords.y
	for isUnoccupied {
		_, isUnoccupied = discover(point{con.activeSandCoords.x, newY - 1}, s)
		if isUnoccupied {
			newY--
		}
	}

	//check new position for sand
	//1 down
	oldCoords := con.activeSandCoords
	newSandCoord := con.activeSandCoords
	if moved == false {
		newSandCoord.y++
		_, occupied := discover(newSandCoord, s)
		if !occupied {
			insertPoint(newSandCoord, s, activeSand)
			con.activeSandCoords = newSandCoord
			moved = true
		}
	}
	if moved == false {
		newSandCoord.x--
		_, occupied := discover(newSandCoord, s)
		if !occupied {
			insertPoint(newSandCoord, s, activeSand)
			con.activeSandCoords = newSandCoord
			moved = true
		}
	}
	if moved == false {
		newSandCoord.x = newSandCoord.x + 2
		_, occupied := discover(newSandCoord, s)
		if !occupied {
			insertPoint(newSandCoord, s, activeSand)
			con.activeSandCoords = newSandCoord
			moved = true
		}

	}
	//unset sand from old position
	if moved {
		delete(s[oldCoords.y], oldCoords.x)
		if !con.infFloor {
			//check for dropout
			if newSandCoord.y >= con.lowestRockPos {
				delete(s[newSandCoord.y], newSandCoord.y)
				return true
			}
		} else {
			if newSandCoord.y == con.infFloorPos-1 {
				insertPoint(newSandCoord, s, sand)
				con.activeSand = false
			}
		}
	} else {
		//inactivate
		insertPoint(oldCoords, s, sand)
		con.activeSand = false
	}

	return false
}

func findLowestRockPos(s map[int]map[int]rune) int {
	lowestRockPos := math.MinInt
	for y, yv := range s {
		for x, _ := range yv {
			if s[y][x] == rock {
				lowestRockPos = utils.Max(lowestRockPos, y)
			}
		}
	}

	return lowestRockPos
}

func spawnSand(con *helperContainer, s map[int]map[int]rune) {
	if !con.activeSand {
		insertPoint(con.entryPoint, s, activeSand)
		con.activeSandCoords = con.entryPoint
		con.activeSand = true
	}
}

func discover(p point, s map[int]map[int]rune) (rune, bool) {
	char, ok := s[p.y][p.x]
	if !ok {
		return '0', false
	}

	return char, true
}

func insertPoint(p point, s map[int]map[int]rune, r rune) {
	checkSurface(p.y, s)
	s[p.y][p.x] = r
}

func insertRocks(command string, surface map[int]map[int]rune) {

	p1, p2 := point{0, 0}, point{0, 0}
	coordinates := strings.Split(command, " -> ")
	for i, coord := range coordinates {
		p1 = p2

		x, _ := strconv.Atoi(coord[:strings.Index(coord, ",")])
		y, _ := strconv.Atoi(coord[strings.Index(coord, ",")+1:])

		p2 = point{x: x, y: y}
		if i == 0 {
			continue
		}

		//draw horizontal
		if p1.y == p2.y {
			for x := utils.Min(p1.x, p2.x); x <= utils.Max(p1.x, p2.x); x++ {
				insertPoint(point{x, p1.y}, surface, rock)
			}
		} else {
			//draw vertical
			for y := utils.Min(p1.y, p2.y); y <= utils.Max(p1.y, p2.y); y++ {
				insertPoint(point{p1.x, y}, surface, rock)
			}
		}
	}

}

func checkSurface(y int, surface map[int]map[int]rune) {
	_, ok := surface[y]
	if !ok {
		surface[y] = make(map[int]rune)
	}
}

func print(s map[int]map[int]rune) {
	clrScr()
	topLeft, bottomRight := borderPoints(s)
	for y := topLeft.y; y <= bottomRight.y; y++ {
		fmt.Println("")
		for x := topLeft.x; x <= bottomRight.x; x++ {
			v, ok := s[y][x]
			if ok {
				fmt.Print(string(v))
			} else {
				fmt.Print(" ")
			}
		}
	}
}

func borderPoints(s map[int]map[int]rune) (point, point) {
	topLeft, bottomRight := point{
		math.MaxInt,
		math.MaxInt,
	}, point{
		math.MinInt,
		math.MinInt,
	}
	for y, yv := range s {
		for x, _ := range yv {
			topLeft.x = utils.Min(topLeft.x, x)
			topLeft.y = utils.Min(topLeft.y, y)
			bottomRight.x = utils.Max(bottomRight.x, x)
			bottomRight.y = utils.Max(bottomRight.y, y)
		}
	}

	return topLeft, bottomRight
}

func clrScr() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

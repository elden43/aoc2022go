package day15

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"github.com/elden43/aoc2022/utils"
	"regexp"
)

type sensorBeaconPair struct{
	sensor utils.Coord
	beacon utils.Coord
	md int
}

var inputPairs []*sensorBeaconPair

func Part1() {
	resLine := 2000000
	field := utils.NewField2DRune()
	inputPairs = processInput()
	for _, pair := range inputPairs {
		drawNaughtyPicture(*pair, field, resLine, resLine)
		field.Add(pair.sensor, 'S')
		field.Add(pair.beacon, 'B')
	}

	fmt.Println(countCharsOnYPos(resLine, '#', field))
}

func Part2() {
	areaCoord1 := utils.Coord{X: 0, Y: 0}
	areaCoord2 := utils.Coord{X: 4000000, Y: 4000000}
	inputPairs = processInput()

	for y:= areaCoord1.Y; y <= areaCoord2.Y; y++ {
		for x:= areaCoord1.X; x <= areaCoord2.X; {
			xChanged := false
			for _, pair := range inputPairs {
				if pair.md >= manhattanDist4Coords(utils.Coord{X: x, Y: y}, pair.sensor) {
					x = pair.sensor.X + pair.md - utils.Abs(y - pair.sensor.Y)+1
					xChanged = true
				}
			}

			if !xChanged {
				println(x * 4000000 + y)
				return
			}

		}
	}
}

func countCharsOnYPos(yPos int, char rune, f *utils.Field2DRune) int {
	res := 0
	for _, v := range f.Field[yPos] {
		if v == char {
			res++
		}
	}

	return res
}

func drawNaughtyPicture(sbp sensorBeaconPair, f *utils.Field2DRune, minLine, maxLine int) {
	//manhattan distance
	//if your speed is 1 manhattan distance per day, md also means "mandays needed to get from start to destination"
	md := manhattanDist4Pair(sbp)

	//if picture wouldn't cross our result line, skip drawing
	if maxLine < sbp.sensor.Y-md || minLine > sbp.sensor.Y+md {
		return
	}

	//draw only in defined range -- yeah, this was line until I started
	//optimize it for part 2 ant before I realized it will not be possible
	minY := sbp.sensor.Y - md
	maxY := sbp.sensor.Y + md
	for y := minY; y <= maxY; y++ {
		if y >= minLine && y <= maxLine {
			yDiff := sbp.sensor.Y - y
			xDiff := utils.Abs(utils.Abs(yDiff) - md)
			for x := sbp.sensor.X - xDiff; x <= sbp.sensor.X+xDiff; x++ {
				f.Add(utils.Coord{X: x, Y: y}, '#')
			}
		}
	}
}

func manhattanDist4Pair(sbp sensorBeaconPair) int {
	return utils.Abs(sbp.sensor.X-sbp.beacon.X) + utils.Abs(sbp.sensor.Y-sbp.beacon.Y)
}

func manhattanDist4Coords(first, second utils.Coord) int {
	return utils.Abs(first.X-second.X) + utils.Abs(first.Y-second.Y)
}

func processInput() []*sensorBeaconPair {
	var res []*sensorBeaconPair
	input := reader.GetInputSeparatedBy("puzzles/day15/input.txt", "\r\n")
	for _, s := range input {
		re := regexp.MustCompile(`(?m)^Sensor at x=(-?[[:digit:]]*), y=(-?[[:digit:]]*): closest beacon is at x=(-?[[:digit:]]*), y=(-?[[:digit:]]*)$`)
		match := re.FindStringSubmatch(s)
		sCoord := utils.Coord{X: utils.Str2int(match[1]), Y: utils.Str2int(match[2])}
		bCoord := utils.Coord{X: utils.Str2int(match[3]), Y: utils.Str2int(match[4])}
		res = append(
			res,
			&sensorBeaconPair{
				sensor: sCoord,
				beacon: bCoord,
				md: manhattanDist4Coords(sCoord, bCoord),
			},
		)
	}

	return res
}

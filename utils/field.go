package utils

import (
	"fmt"
	"math"
)

type Field2DRune struct {
	Field map[int]map[int]rune
}

func NewField2DRune() *Field2DRune {
	m := make(map[int]map[int]rune)
	return &Field2DRune{m}
}

//control methods

func (f *Field2DRune) Add(c Coord, r rune) {
	f.check(c)
	f.Field[c.Y][c.X] = r
}

// debug&print methods

func (f *Field2DRune) Print(emptyChar rune) {
	topLeft, bottomRight := f.borderCoords()
	for y := topLeft.Y; y <= bottomRight.Y; y++ {
		fmt.Println("")
		for x := topLeft.X; x <= bottomRight.X; x++ {
			v, ok := f.Field[y][x]
			if ok {
				fmt.Print(string(v))
			} else {
				fmt.Print(string(emptyChar))
			}
		}
	}
}


//internal methods

func (f *Field2DRune) check(c Coord) {
	_, ok := f.Field[c.Y]
	if !ok {
		f.Field[c.Y] = make(map[int]rune)
	}
}

func (f *Field2DRune) borderCoords() (topLeft Coord, bottomRight Coord) {
	topLeft, bottomRight = Coord{
		math.MaxInt,
		math.MaxInt,
	}, Coord{
		math.MinInt,
		math.MinInt,
	}
	for y, yv := range f.Field {
		for x, _ := range yv {
			topLeft.X = Min(topLeft.X, x)
			topLeft.Y = Min(topLeft.Y, y)
			bottomRight.X = Max(bottomRight.X, x)
			bottomRight.Y = Max(bottomRight.Y, y)
		}
	}

	return
}


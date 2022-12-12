package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"math"
	"strconv"
	"strings"
)

type file struct {
	name string
	path string
	size uint32
}

func Day7part1() {
	input := reader.GetInputSeparatedBy("puzzles/day7/input.txt", "\n")
	files := prepFiles(input)
	dirSizes := prepDirs(input, files)

	size := uint32(0)
	for _, v := range dirSizes {
		if v < 100000 {
			size += v
		}
	}

	fmt.Println(size)
}

func Day7part2() {
	fsTotal := uint32(70000000)
	fsNeeded := uint32(30000000)

	input := reader.GetInputSeparatedBy("puzzles/day7/input.txt", "\n")
	files := prepFiles(input)
	dirSizes := prepDirs(input, files)

	smallestSize := uint32(math.MaxUint32)
	sizeToClean := fsNeeded - (fsTotal - dirSizes["~/"])
	for _, v := range dirSizes {
		if v > sizeToClean && v < smallestSize {
			smallestSize = v
		}
	}

	fmt.Println(smallestSize)
}

func prepDirs(input []string, files []*file) map[string]uint32 {
	dirSizes := make(map[string]uint32)
	currentPath := ""

	for _, v := range input {
		if len(v) >= 7 && v[:7] == "$ cd .." {
			currentPath = currentPath[:strings.LastIndex(currentPath, "~")]
		} else if len(v) >= 5 && v[:5] == "$ cd " {
			currentPath += "~" + v[5:len(v)-1]
			for _, fv := range files {
				_, ok := dirSizes[currentPath]
				if !ok {
					dirSizes[currentPath] = uint32(0)
				}
				if len(fv.path) >= len(currentPath) {
				}
				if len(fv.path) >= len(currentPath) && fv.path[:len(currentPath)] == currentPath {
					dirSizes[currentPath] += fv.size
				}

			}
		}
	}
	return dirSizes
}

func prepFiles(input []string) []*file {
	var files []*file
	var currentPath string
	for _, v := range input {
		if len(v) >= 7 && v[:7] == "$ cd .." {
			currentPath = currentPath[:strings.LastIndex(currentPath, "~")]
		} else if len(v) >= 5 && v[:5] == "$ cd " {
			currentPath += "~" + v[5:len(v)-1]
		} else if v[:1] != "$" && v[:4] != "dir " {
			u64, _ := strconv.ParseUint(v[:strings.Index(v, " ")], 10, 32)
			size := uint32(u64)
			file := file{v[strings.Index(v, " ") : len(v)-1], currentPath, size}
			files = append(files, &file)
		}
	}
	return files
}

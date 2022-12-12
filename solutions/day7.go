package solutions

import (
	"fmt"
	"github.com/elden43/aoc2022/reader"
	"math"
	"strconv"
	"strings"
)

//type fileSystemType int
//
//const (
//	file fileSystemType = iota
//	dir
//)
//
//type fileSystemRecord struct {
//	path           *filesystemPath
//	fileSystemType fileSystemType
//	name           string
//	size           uint32
//}
//
//type command struct {
//	command string
//	results []string
//}
//
//func (c command) printlnResults() {
//	for _, v := range c.results {
//		fmt.Println(v)
//	}
//}
//
//type filesystemPath struct {
//	path map[int]string
//}
//
//func (fp filesystemPath) add(dir string) {
//	fp.path[len(fp.path)] = dir
//}
//
//func (fp filesystemPath) removeLast() {
//	delete(fp.path, len(fp.path)-1)
//}
//
//func (fp filesystemPath) printPath() {
//	for i := 0; i < len(fp.path); i++ {
//		fmt.Print(fp.path[i])
//		fmt.Print("~")
//	}
//	fmt.Println("")
//}
//
//func newFilesystemPath() *filesystemPath {
//	fsp := new(filesystemPath)
//	fsp.path = make(map[int]string)
//	return fsp
//}
//
//func Day7part1() {
//	commands := parseCommands("puzzles/day7/sample.txt")
//	var fs []fileSystemRecord
//
//	currentPath := newFilesystemPath()
//	for i := 0; i < len(commands); i++ {
//		if strings.Contains(commands[i].command, "cd ..") {
//			currentPath.removeLast()
//		} else if strings.Contains(commands[i].command, "cd") {
//			currentPath.add(commands[i].command[3 : len(commands[i].command)-1])
//			currentPath.printPath()
//		} else if strings.Contains(commands[i].command, "ls") {
//			for _, v := range commands[i].results {
//				if strings.Contains(v, "dir") {
//					cp := currentPath
//					fsr := fileSystemRecord{cp, dir, v[strings.LastIndex(v, " ") : len(v)-1], 0}
//					fs = append(fs, fsr)
//				} else {
//					ui64, _ := strconv.ParseUint(v[:strings.LastIndex(v, " ")], 10, 64)
//					fsr := fileSystemRecord{currentPath, file, v[strings.LastIndex(v, " ") : len(v)-1], uint32(ui64)}
//					fs = append(fs, fsr)
//				}
//			}
//		}
//	}
//
//	fmt.Println(fs)
//}
//
//func Day7part2() {
//	/*for _, line := range reader.ReadLines("puzzles/day6/input.txt") {
//
//	}*/
//	fmt.Println("not yet")
//}
//
//func parseCommands(file string) map[int]command {
//	input := reader.GetInputSeparatedBy("puzzles/day7/sample.txt", "\n")
//
//	commands := make(map[int]command)
//	currentCommand := ""
//	var currentResults []string
//	iC := 0
//	for i, line := range input {
//		if strings.Contains(line, "$") || i == len(input)-1 {
//			if currentCommand == "" {
//				currentCommand = line[2:]
//			} else {
//				command := command{currentCommand, currentResults}
//				currentResults = nil
//				commands[iC] = command
//				iC++
//				currentCommand = line[2:]
//			}
//		} else {
//			currentResults = append(currentResults, line)
//		}
//	}
//
//	/*for i := 0; i < len(commands); i++ {
//		fmt.Println(commands[i].command)
//		commands[i].printlnResults()
//	}*/
//	return commands
//}

type file struct {
	name string
	path string
	size uint32
}

func Day7part1() {
	var files []*file
	var currentPath string

	input := reader.GetInputSeparatedBy("puzzles/day7/input.txt", "\n")
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

	dirSizes := make(map[string]uint32)
	currentPath = ""

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

	size := uint32(0)
	for _, v := range dirSizes {
		if v < 100000 {
			size += v
		}
	}

	fmt.Println(size)
}

func Day7part2() {
	var files []*file
	var currentPath string
	fsTotal := uint32(70000000)
	fsNeeded := uint32(30000000)

	input := reader.GetInputSeparatedBy("puzzles/day7/input.txt", "\n")
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

	dirSizes := make(map[string]uint32)
	currentPath = ""

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

	smallestSize := uint32(math.MaxUint32)
	sizeToClean := fsNeeded - (fsTotal - dirSizes["~/"])
	for _, v := range dirSizes {
		if v > sizeToClean && v < smallestSize {
			smallestSize = v
		}
	}

	fmt.Println(smallestSize)
}

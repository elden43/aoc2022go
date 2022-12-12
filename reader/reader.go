package reader

import (
	"bufio"
	"os"
	"strings"
)

func ReadLines(filePath string) []string {
	input, _ := os.Open(filePath)
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	scanner := bufio.NewScanner(input)

	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func GetInputSeparatedBy(file, separator string) []string {
	bytes, _ := os.ReadFile(file)
	return strings.Split(string(bytes), separator)
}

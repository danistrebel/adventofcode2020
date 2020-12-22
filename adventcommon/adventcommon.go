package adventcommon

import (
	"os"
	"bufio"
)

func ParseInput(filePath string) []string  {
	file, err := os.Open(filePath)


	if err != nil {
		os.Exit(-1)
	}

	defer file.Close()

	var allInput []string


	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		allInput = append(allInput, scanner.Text())
	}
	return allInput
}

func ParseCharMatrix(filePath string) [][]rune  {
	var charMatrix [][]rune

	for _, row := range ParseInput(filePath) {
		var charRow []rune
		for _, c := range row {
			charRow = append(charRow, c)
		}
		charMatrix = append(charMatrix, charRow)
	}

	return charMatrix
}
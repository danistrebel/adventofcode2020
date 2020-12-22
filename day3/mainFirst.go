package main

import (
	"../adventcommon"
	"fmt"
)

func main() {
	
	inputMatrix := adventcommon.ParseCharMatrix("input.txt")
	fmt.Println(inputMatrix)

	directionRight := 3
	currentX := 0
	
	widthPattern := len(inputMatrix[0])
	treeChar := []rune("#")[0]
	treeCount := 0

	for currentY := 0; currentY < len(inputMatrix); currentY++ {
		slopeTile := inputMatrix[currentY][currentX]

		fmt.Println(currentY,currentX, slopeTile)

		if (slopeTile == treeChar) {
			treeCount++
		}
		currentX = (currentX+directionRight) % widthPattern

	}

	fmt.Printf("\nTrees: %d", treeCount)
}
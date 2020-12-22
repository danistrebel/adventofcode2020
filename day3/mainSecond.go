package main

import (
	"../adventcommon"
	"fmt"
)

type skiInstr struct {
	down int
	right int
}

func main() {
	
	inputMatrix := adventcommon.ParseCharMatrix("input.txt")
	// fmt.Println(inputMatrix)
	widthPattern := len(inputMatrix[0])
	treeChar := []rune("#")[0]

	instructions := []skiInstr{
		skiInstr{down: 1, right: 1},
		skiInstr{down: 1, right: 3},
		skiInstr{down: 1, right: 5},
		skiInstr{down: 1, right: 7},
		skiInstr{down: 2, right: 1},
	}

	treeProduct := 1

	for _, instr := range instructions {
		// fmt.Println(instr)
		directionDown := instr.down
		directionRight:= instr.right
		fmt.Printf("\n(down %d, right %d)", directionDown, directionRight)

		currentX := 0
		
		treeCount := 0

		for currentY := 0; currentY < len(inputMatrix); currentY+=directionDown {
			slopeTile := inputMatrix[currentY][currentX]

			// fmt.Println(currentY,currentX, slopeTile)

			if (slopeTile == treeChar) {
				treeCount++
			}
			currentX = (currentX+directionRight) % widthPattern

		}

		fmt.Printf("\n(%d,%d)Trees: %d", directionDown, directionRight, treeCount)
		treeProduct *= treeCount
	}

	fmt.Printf("\nTree Product %d", treeProduct)
}
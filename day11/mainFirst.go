package main

import (
	"../adventcommon"
	"fmt"
)

const emptySeat = 'L'
const occupiedSeat = '#'
const floor = '.'


func countOccupiedNeighbors(matrix [][]rune, row int, col int) int {
	occupiedNeighbors := 0
	if row-1 >= 0 {
		if col-1 >= 0 && matrix[row-1][col-1] == occupiedSeat { occupiedNeighbors++ }
		if matrix[row-1][col] == occupiedSeat { occupiedNeighbors++ }
		if col+1 < len(matrix[row-1]) && matrix[row-1][col+1] == occupiedSeat { occupiedNeighbors++ }
	}
	if col-1 >= 0 && matrix[row][col-1] == occupiedSeat { occupiedNeighbors++ }
	if col+1 < len(matrix[row]) && matrix[row][col+1] == occupiedSeat { occupiedNeighbors++ }
	if row+1 < len(matrix) {
		if col-1 >= 0 && matrix[row+1][col-1] == occupiedSeat { occupiedNeighbors++ }
		if matrix[row+1][col] == occupiedSeat { occupiedNeighbors++ }
		if col+1 < len(matrix[row+1]) && matrix[row+1][col+1] == occupiedSeat { occupiedNeighbors++ }
	}
	return occupiedNeighbors
}

func countOccupiedSeats(matrix [][]rune) int  {
	count := 0
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == occupiedSeat {
				count++
			}
		}
	}
	return count
}

func main() {
	
	inputMatrix := adventcommon.ParseCharMatrix("input.txt")

	changedInPrevStep := true

	for changedInPrevStep {

		// Deep copy matrix
		nextStepMatrix := make([][]rune, len(inputMatrix))
		copy(nextStepMatrix, inputMatrix)
		for i := range inputMatrix {
			nextStepMatrix[i] = make([]rune, len(inputMatrix[i]))
			copy(nextStepMatrix[i], inputMatrix[i])
		}

		changedInPrevStep = false
		for row := range inputMatrix {
			for col := range inputMatrix[row] {
				if inputMatrix[row][col] == emptySeat && countOccupiedNeighbors(inputMatrix, row,col) == 0 {
					changedInPrevStep = true
					nextStepMatrix[row][col] = occupiedSeat
				} else if inputMatrix[row][col] == occupiedSeat && countOccupiedNeighbors(inputMatrix, row,col) >= 4 {
					changedInPrevStep = true
					nextStepMatrix[row][col] = emptySeat
				}
			}
		}

		inputMatrix = nextStepMatrix
	}

	fmt.Println("final occupied seats: ", countOccupiedSeats(inputMatrix))

}
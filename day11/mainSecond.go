package main

import (
	"../adventcommon"
	"fmt"
)

const emptySeat = 'L'
const occupiedSeat = '#'
const floor = '.'


func countVisibleOccupiedSeats(matrix [][]rune, row int, col int) int {
	occupiedNeighbors := 0

	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1,0}, {1,1},
	}

	for _, direction := range directions {
		dx := direction[0]
		dy := direction[1]

		rowX := row
		colY := col

		for true {
			rowX+=dx
			colY+=dy

			if rowX < 0 || rowX >= len(matrix) || colY < 0 || colY >= len(matrix[rowX]) {
				break
			}

			if matrix[rowX][colY] == occupiedSeat {
				occupiedNeighbors++
				break
			}

			if matrix[rowX][colY] == emptySeat {
				break
			} 
		}
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

	gen := 0

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
				if inputMatrix[row][col] == emptySeat && countVisibleOccupiedSeats(inputMatrix, row,col) == 0 {
					changedInPrevStep = true
					nextStepMatrix[row][col] = occupiedSeat
				} else if inputMatrix[row][col] == occupiedSeat && countVisibleOccupiedSeats(inputMatrix, row,col) >= 5 {
					changedInPrevStep = true
					nextStepMatrix[row][col] = emptySeat
				}
			}
		}
		gen++
		fmt.Println("generation:", gen)

		inputMatrix = nextStepMatrix
	}

	fmt.Println("final occupied seats: ", countOccupiedSeats(inputMatrix))

}
package main

import (
	"../adventcommon"
	"fmt"
)

func calcBinary(input []rune, loMarker rune, hiMarker rune, iLo int, iHi int) int {
	if input[0] == loMarker {
		if len(input) == 1 { return iLo }
		return calcBinary(input[1:], loMarker, hiMarker, iLo, iLo + (iHi - iLo)/2)
	} else if input[0] == hiMarker{
		if len(input) == 1 { return iHi }
		return calcBinary(input[1:], loMarker, hiMarker, iLo + (iHi + 1 - iLo)/2, iHi)
	}
	return 0
}

func main() {
	
	inputMatrix := adventcommon.ParseCharMatrix("input.txt")

	fChar := []rune("F")[0]
	bChar := []rune("B")[0]
	lChar := []rune("L")[0]
	rChar := []rune("R")[0]	

	maxSeatID := -1
	for _, seatCode := range inputMatrix {
		row := calcBinary(seatCode[0:7], fChar, bChar, 0, 127)
		col := calcBinary(seatCode[7:], lChar, rChar, 0, 7)
		seatID := row * 8 + col 
		fmt.Printf("%v: row %d, col: %d, seat: %d\n", seatCode, row, col, seatID)

		if (seatID > maxSeatID) { maxSeatID = seatID }
	}

	fmt.Println("Max Seat ID", maxSeatID)


}
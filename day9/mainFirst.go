package main

import (
	"../adventcommon"
	"fmt"
	"strconv"
)

func main()  {
	allLines := adventcommon.ParseInput("input.txt")

	preambleLength := 25

	inputNumbers := []int{} 

	for i := 0; i < len(allLines); i++ {
		inputNumber32, _ := strconv.ParseInt(allLines[i], 10, 32)
		inputNumbers = append(inputNumbers, int(inputNumber32))
	}

	for i := preambleLength; i < len(inputNumbers); i++ {

		targetSum := inputNumbers[i];
		candiates := make(map[int]bool)
		foundSum := false

		for j := i-preambleLength; j< i; j++ {
			if _, exists := candiates[inputNumbers[j]]; exists {
				foundSum = true
				break;
			}
			candiates[targetSum - inputNumbers[j]] = true
		}
		
		if !foundSum {
			fmt.Println("No sum found for", targetSum)
			break
		}
	}
}
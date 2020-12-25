package main

import (
	"../adventcommon"
	"fmt"
	"strconv"
)

func main()  {
	allLines := adventcommon.ParseInput("input.txt")

	targetSum := 29221323

	inputNumbers := []int{} 

	for i := 0; i < len(allLines); i++ {
		inputNumber32, _ := strconv.ParseInt(allLines[i], 10, 32)
		inputNumbers = append(inputNumbers, int(inputNumber32))
	}


	for i := 0; i < len(inputNumbers); i++ {

		if inputNumbers[i] == targetSum {
			continue
		}

		runningSum := inputNumbers[i];
		min := inputNumbers[i]
		max := inputNumbers[i]

		for j := i+1; runningSum < targetSum ; j++ {
			runningSum += inputNumbers[j]
			if inputNumbers[j] < min { min = inputNumbers[j] }
			if inputNumbers[j] > max { max = inputNumbers[j] }
		}

		if runningSum == targetSum {
			fmt.Println("found range starting at: ", inputNumbers[i], "min:", min, ", max:", max, ", sum:", min+max)
		}
	}
}
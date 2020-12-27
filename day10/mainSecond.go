package main

import (
	"../adventcommon"
	"fmt"
	"strconv"
	"sort"
)

func main()  {
	allLines := adventcommon.ParseInput("input.txt")
	inputNumbers := make([]int, len(allLines)) 
	intMap := make(map[int]bool)

	//always start with 0
	intMap[0] = true

	for i := 0; i < len(allLines); i++ {
		inputNumber32, _ := strconv.ParseInt(allLines[i], 10, 32)
		inputNumbers[i] = int(inputNumber32)
		intMap[int(inputNumber32)] = true
	}

	inputNumbers = append(inputNumbers, 0)
	sort.Ints(inputNumbers)
	maxValue := inputNumbers[len(inputNumbers)-1]
	inputNumbers = append(inputNumbers, maxValue+3)

	combinationsTracker := make([]int, maxValue + 1)
	combinationsTracker[0] = 1 //only one way to start at zero

	for i := 1; i < len(combinationsTracker); i++  {
		for diff:=1; diff < 4; diff++ {
			if _, exists := intMap[i-diff]; exists {
				combinationsTracker[i]+=combinationsTracker[i-diff]
			}
		}
	}

	fmt.Println("combinations:", combinationsTracker[len(combinationsTracker)-1])


}
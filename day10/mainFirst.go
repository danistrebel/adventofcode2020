package main

import (
	"../adventcommon"
	"fmt"
	"strconv"
	"sort"
)

func main()  {
	allLines := adventcommon.ParseInput("input.txt")

	inputNumbers := []int{} 

	for i := 0; i < len(allLines); i++ {
		inputNumber32, _ := strconv.ParseInt(allLines[i], 10, 32)
		inputNumbers = append(inputNumbers, int(inputNumber32))
	}

	sort.Ints(inputNumbers)

	count1 := 0
	count3 := 0

	previousOut := 0

	for i := 0; i < len(inputNumbers); i++  {
		if previousOut+1 == inputNumbers[i] {
			count1++
		} else if previousOut+3 == inputNumbers[i] {
			count3++
		} else {
			fmt.Println("other jump", inputNumbers[i+1]-previousOut)
		}
		previousOut = inputNumbers[i]
	}

	count3++

	fmt.Println("1:", count1, "3:", count3, "=>", count1*count3)

}
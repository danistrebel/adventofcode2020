package main

import (
	"../adventcommon"
	"fmt"
)

func main()  {

	// make sure last entry is counted
	allLines := append(adventcommon.ParseInput("input.txt"), "") 

	totalYes := 0

	groupAnswers := make(map[rune]bool)

	for _, line := range allLines {
		if len(line) == 0 {
			yesInGroup:=0
			for range groupAnswers {
				yesInGroup++
			}
			totalYes += yesInGroup
			groupAnswers = make(map[rune]bool)
		} else {
			yesPerPerson := []rune(line)
			for _, questionID := range yesPerPerson {
				groupAnswers[questionID] = true
			}
		}
	}

	fmt.Println("Total Yes:", totalYes)

}
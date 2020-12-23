package main

import (
	"../adventcommon"
	"fmt"
)

func main()  {

	// make sure last entry is counted
	allLines := append(adventcommon.ParseInput("input.txt"), "") 

	totalYes := 0

	groupAnswers := make(map[rune]int)
	groupSize := 0

	for _, line := range allLines {
		if len(line) == 0 {
			yesInGroup:=0
			for _, val := range groupAnswers {
				if (val == groupSize) {
					yesInGroup++
				}
			}
			totalYes += yesInGroup
			groupAnswers = make(map[rune]int)
			groupSize=0
		} else {
			yesPerPerson := []rune(line)
			for _, questionID := range yesPerPerson {
				if val, exists := groupAnswers[questionID]; exists {
					groupAnswers[questionID] = val+1
				} else{
					groupAnswers[questionID] = 1
				}
			}
			groupSize++
		}
	}

	fmt.Println("Total Yes:", totalYes)

}
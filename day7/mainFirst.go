package main

import (
	"../adventcommon"
	"fmt"
	"regexp"
	"strings"
)

func main()  {
	allLines := adventcommon.ParseInput("input.txt")

	parentBags := make(map[string][]string)

	rOuter := regexp.MustCompile(`(\w+\s\w+) bags? contain (.*)`)
	rInner := regexp.MustCompile(`\d+ (\w+\s\w+) bags?`)

	for _, line := range allLines {		
		fmt.Println(line)
		matches := rOuter.FindStringSubmatch(line)

		outerBag := matches[1]
		for _, match := range matches[2:] {

			for _, innerBagEntry := range strings.Split(string(match), ",") {
				innerMatches := rInner.FindStringSubmatch(innerBagEntry)
				if len(innerMatches) > 1 {
					innerBag := innerMatches[1]
					fmt.Println("--", innerBag, "(", outerBag, ")")

					if entry, exists := parentBags[innerBag]; exists {
						parentBags[innerBag] = append(entry, outerBag)
					} else {
						parentBags[innerBag] = []string{outerBag}
					}
				}
			}
		}
	}	

	bagsForShinyGold := make(map[string]bool)

	bagsToCheck := parentBags["shiny gold"]

	for len(bagsToCheck) > 0 {
		bagToCheck := bagsToCheck[0]
		bagsToCheck = bagsToCheck[1:]

		if _, exists := bagsForShinyGold[bagToCheck]; exists {
			// to prevent loops
			continue
		}

		if parentsToCheck, exists := parentBags[bagToCheck]; exists {
			bagsToCheck = append(bagsToCheck, parentsToCheck...)
		}

		bagsForShinyGold[bagToCheck] = true

	}

	fmt.Println(bagsForShinyGold)

	fmt.Println("Total possible bags: ", len(bagsForShinyGold))
}
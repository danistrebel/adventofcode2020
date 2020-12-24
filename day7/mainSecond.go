package main

import (
	"../adventcommon"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

type bagContent struct {
	count int
	color string
}

func main()  {
	allLines := adventcommon.ParseInput("input.txt")

	rOuter := regexp.MustCompile(`(\w+\s\w+) bags? contain (.*)`)
	rInner := regexp.MustCompile(`(\d+) (\w+\s\w+) bags?`)

	bagContentMap := make(map[string][]bagContent)

	for _, line := range allLines {		
		fmt.Println(line)
		matches := rOuter.FindStringSubmatch(line)

		outerBag := matches[1]
		innerBags := []bagContent{}
		for _, match := range matches[2:] {

			for _, innerBagEntry := range strings.Split(string(match), ",") {
				innerMatches := rInner.FindStringSubmatch(innerBagEntry)
				if len(innerMatches) > 1 {

					innerBagCount, _ := strconv.ParseInt(innerMatches[1], 10, 32)
					innerBagColor := innerMatches[2]
					fmt.Println("--", outerBag, "(", innerBagCount, "," , innerBagColor , ")")

					innerBags = append(innerBags, bagContent{count: int(innerBagCount), color: innerBagColor})
				}
			}
		}

		bagContentMap[outerBag]=innerBags

	}	

	totalCount := countInnerBags("shiny gold", bagContentMap)
	totalCount-- // not counting the shiny gold one

	fmt.Println("total bags in shiny gold:", totalCount)
}

func countInnerBags(bagToAdd string, bagContentMap map[string][]bagContent) int {
	count := 1
	for _, bc := range bagContentMap[bagToAdd] {
		fmt.Println(bc)
		count += (bc.count * countInnerBags(bc.color, bagContentMap))
	}
	fmt.Println(bagToAdd, count)
	return count

}
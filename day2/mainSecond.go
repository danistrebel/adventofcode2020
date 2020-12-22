package main

import (
	"../adventcommon"
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	
	r := regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]+)`)
	correctCount:=0

	for _, passInput := range adventcommon.ParseInput("input.txt") {
		matches := r.FindStringSubmatch(passInput)
		fmt.Println(passInput)
		
		splits := strings.Split(matches[4], "")

		first, _ := strconv.ParseInt(matches[1], 10, 32)
		second, _ := strconv.ParseInt(matches[2], 10, 32)

		fistTrue := splits[int(first)-1] == matches[3]
		secondTrue := splits[int(second)-1] == matches[3]

		if fistTrue!=secondTrue {
			correctCount++
			fmt.Println(" correct!")
		}
	}

	fmt.Printf("Correctly formatted: %d", correctCount)
}
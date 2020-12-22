package main

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	r := regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]+)`)

	correctCount:=0
	
	for scanner.Scan() {

		passInput := scanner.Text()
		matches := r.FindStringSubmatch(passInput)
		// fmt.Printf("%#v\n", )
		fmt.Println(passInput)

		charCount := 0
		
		for _,char := range strings.Split(matches[4], "") {
			if (char == matches[3]) {
				charCount++
			}
		} 

		minCharCount, _ := strconv.ParseInt(matches[1], 10, 32)
		maxCharCount, _ := strconv.ParseInt(matches[2], 10, 32)

		if charCount >= int(minCharCount) && charCount <= int(maxCharCount) {
			correctCount++
			fmt.Println(" correct!")
		}
	}

	fmt.Printf("Correctly formatted: %d", correctCount)
}
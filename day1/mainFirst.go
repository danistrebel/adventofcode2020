package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		os.Exit(-1)
	}

	defer file.Close()


	scanner := bufio.NewScanner(file)
	targetSum := 2020

	seen := make(map[int]int)

	for scanner.Scan() {
		i64, err := strconv.ParseInt(scanner.Text(), 10, 32)

		i:=int(i64)

		if err != nil {
			fmt.Println("Parsing error: %d", err)
			os.Exit(-1);
		}

		if prev, ok := seen[i]; ok {
			fmt.Printf("prev: %d, curr: %d, multiplied: %d\n", prev, i, prev*i)
			break;
		}

		seen[targetSum-i] = i
	}

	fmt.Println("done")
}
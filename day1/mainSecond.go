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

	allNumbers := make(map[int]bool)

	for scanner.Scan() {
		i64, err := strconv.ParseInt(scanner.Text(), 10, 32)

		i:=int(i64)

		if err != nil {
			fmt.Println("Parsing error: %d", err)
			os.Exit(-1);
		}

		allNumbers[i] = true
	}

	for k1 := range allNumbers {
		for k2 := range allNumbers {
			if k1 != k2 {
				k3 := targetSum - k1 - k2
				if _, found := allNumbers[k3]; found {
					fmt.Printf("1: %d, 2: %d, 3: %d, product: %d", k1, k2, k3, k1*k2*k3)
					os.Exit(0);
				}
			}
		}
	}

	fmt.Println("done")
}
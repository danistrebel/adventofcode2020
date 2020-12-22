package main

import (
	"../adventcommon"
	"fmt"
	"strings"
)

func main()  {

	currentPassfields := make(map[string]string) 
	requiredFields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid", Optional!
	}

	validPassports := 0
	passportCount := 0 

	// make sure last entry is counted

	allLines := append(adventcommon.ParseInput("input.txt"), "") 

	for _, passInput := range allLines {
		if len(passInput) == 0 {

			passportIsValid := true

			for _, requiredField := range requiredFields {
				if _, contained := currentPassfields[requiredField]; !contained {
					passportIsValid = false
					break
				}

			}

			if passportIsValid {
				validPassports++
				fmt.Printf("++")
			}

			fmt.Println("pass", currentPassfields)

			passportCount++

			currentPassfields = make(map[string]string) 

		} else {
			for _, passEntry := range strings.Split(passInput, " ") {
				passEntrySplit := strings.Split(passEntry, ":")
				currentPassfields[passEntrySplit[0]] = passEntrySplit[1]
			}
		}
	}

	fmt.Printf("Valid Passports: %d/%d", validPassports, passportCount)
}
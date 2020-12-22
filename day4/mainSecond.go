package main

import (
	"../adventcommon"
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

// field not nil and min <= year <= max
func isInRange(field string, min int, max int) bool {
	fieldInt64, err := strconv.ParseInt(field, 10, 32) 
	if err != nil {
		return false
	}
	fieldInt := int(fieldInt64)
	return fieldInt >= min && fieldInt <= max
}

func isValidHeight(height string) bool {
	if len(height)>3 {
		unit:= height[len(height)-2:len(height)]
		
		if unit == "cm" {
			return isInRange(height[:len(height)-2], 150, 193)
		} else if (unit == "in") {
			return isInRange(height[:len(height)-2], 59, 76)	
		}
		return false
	}
	return false
}

func isValidHairColor(hair string) bool {
	re := regexp.MustCompile(`^#[0-9a-z]{6}$`)
	return re.MatchString(hair)
}

func isValidEyeColor(ecl string) bool {
	switch ecl {
		case "amb":
			return true
		case "blu":
			return true
		case "brn":
			return true
		case "gry":
			return true
		case "grn":
			return true
		case "hzl":
			return true
		case "oth":
			return true
		default:
			return false
	}

	return false
}

func isValidPassportID(pid string) bool {
	re := regexp.MustCompile(`^[0-9]{9}$`)
	return re.MatchString(pid)
}

func isValidPassport(passport map[string]string) bool {
	return isInRange(passport["byr"], 1920, 2002) &&
	isInRange(passport["iyr"], 2010, 2020) &&
	isInRange(passport["eyr"], 2020, 2030) &&
	isValidHeight(passport["hgt"]) &&
	isValidHairColor(passport["hcl"]) &&
	isValidEyeColor(passport["ecl"]) &&
	isValidPassportID(passport["pid"])

}

func main()  {

	currentPassfields := make(map[string]string) 

	validPassports := 0
	passportCount := 0 

	// make sure last entry is counted

	allLines := append(adventcommon.ParseInput("input.txt"), "") 

	for _, passInput := range allLines {
		if len(passInput) == 0 {

			if isValidPassport(currentPassfields) {
				validPassports++
				fmt.Println("[valid] ", currentPassfields)
			}


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
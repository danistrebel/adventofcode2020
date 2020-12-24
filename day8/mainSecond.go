package main

import (
	"../adventcommon"
	"fmt"
	"strconv"
	"reflect"
)

func main()  {
	originalLines := adventcommon.ParseInput("input.txt")
    fmt.Println(reflect.TypeOf(originalLines))

	for instrIndex := range originalLines {
		allLines := []string{}
		allLines = append(allLines, originalLines...)

		//flip single instruction
		if  allLines[instrIndex][0:3] == "jmp" {
			allLines[instrIndex] = "nop"+allLines[instrIndex][3:]
		} else if allLines[instrIndex][0:3] == "nop" {
			allLines[instrIndex] = "jmp"+allLines[instrIndex][3:]
		}

		acc := 0

		for i:= 0; len(allLines[i])!=0; {
			instr := allLines[i]
			opr := instr[0:3]
			sig := instr[4:5]
			val := instr[5:]
			allLines[i]=""
	
			//fmt.Println(opr, sig, val)
			int32Val, _ := strconv.ParseInt(val, 10, 32)
			intVal := int(int32Val)
	
			switch opr {
				case "nop":
					i++
				case "jmp":
					if sig == "+" {
						i+=intVal
					} else if sig == "-" {
						i-=intVal
					} else {
						fmt.Println("ERROR: What is sig:", sig)
					}
				case "acc":
					if sig == "+" {
						acc+=intVal
					} else if sig == "-" {
						acc-=intVal
					} else {
						fmt.Println("ERROR: What is sig:", sig)
					}
					i++
			}

			if i==len(allLines) {
				fmt.Println("REACHED THE END: final acc:", acc)
				break
			}
		} 

	}

	

}
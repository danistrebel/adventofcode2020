package main

import (
	"../adventcommon"
	"fmt"
	"strconv"
)

func main()  {
	allLines := adventcommon.ParseInput("input.txt")

	acc := 0

	for i:= 0; len(allLines[i])!=0; {
		instr := allLines[i]
		opr := instr[0:3]
		sig := instr[4:5]
		val := instr[5:]
		allLines[i]=""

		fmt.Println(opr, sig, val)
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
	} 

	fmt.Println("final acc:", acc)
}
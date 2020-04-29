package main

import "fmt"

func printAsm() {
	var currentByte int

	i := 0
	currentByte = obj[i]

	for currentByte != RET {
		switch currentByte {
		case PUSH:
			i += 1
			fmt.Printf("PUSH %d\n", obj[i])
		case READ:
			i += 1
			currentByte = obj[i]
			fmt.Printf("READ %s\n", getSym(currentByte).name)
		case WRITE:
			i += 1
			currentByte = obj[i]
			fmt.Printf("WRITE %s\n", getSym(currentByte).name)
		case ADD:
			fmt.Println("ADD POP, POP")
		case SUB:
			fmt.Println("SUB POP, POP")
		case MUL:
			fmt.Println("MUL POP, POP")
		case DIV:
			fmt.Println("DIV POP, POP\n")
		}
		i += 1
		currentByte = obj[i]
	}
	fmt.Print("RET \n")
}

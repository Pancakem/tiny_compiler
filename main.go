package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	_ = iota
	SOURCE_ARG
	MIN_ARGS
)

func main() {
	argv := os.Args[0:]
	argc := len(os.Args)

	if argc < MIN_ARGS {
		log.Fatal("No source provided!\nUsage: compiler source")
	}

	file, err := os.Open(argv[1])
	if err != nil {
		log.Fatal(err)
	}

	source = bufio.NewReader(file)

	tok = nextTok()
	compile(parse())

	fmt.Println("Generated ASM:\n")
	printAsm()

	fmt.Println("Execution result:\n")
	run()
}

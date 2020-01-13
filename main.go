package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	filebuffer, err := ioutil.ReadFile(argv[SOURCE_ARG])
	if err != nil {
		log.Fatal(err)
	}

	source = bufio.NewScanner(strings.NewReader(string(filebuffer)))
	source.Split(bufio.ScanRunes)

	tok = nextTok()
	compile(parse())

	fmt.Println("Generated ASM:\n")
	printAsm()

	fmt.Println("Execution result:\n")
	run()
}

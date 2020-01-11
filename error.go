package main

import (
	"fmt"
	"os"
)

func fatalError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
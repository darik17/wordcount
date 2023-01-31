package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if args[0] == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(strings.Split(args[0], " ")))
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		fmt.Println(len(strings.Split(args[0], " ")))
	} else {
		fmt.Println(0)
	}
}

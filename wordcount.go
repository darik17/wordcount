package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(len(strings.Split(args[0], " ")))
}

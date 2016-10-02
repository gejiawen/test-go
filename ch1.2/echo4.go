package main

import (
	"fmt"
	"strings"
)

func main() {
	// TODO
	// use time package to test the time cost.
}

func echo1(args []string) {
	s, sep := "", ""

	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""

	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}

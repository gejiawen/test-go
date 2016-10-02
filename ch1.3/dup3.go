package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		content, err := ioutil.ReadFile(fileName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(content), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s\n", n, line)
		}

	}
}

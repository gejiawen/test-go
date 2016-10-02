package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		// 如果没有文件参数，则从标准输入读取数据流
		countLines(os.Stdin, counts)
		// fmt.Print("file argument required")
	} else {
		for _, file := range files {
			f, err := os.Open(file)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s\n", n, line)
		}
	}
}

func countLines(f *os.File, pool map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		pool[input.Text()]++
	}
}

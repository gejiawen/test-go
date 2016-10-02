package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		ok := strings.HasPrefix(url, "http://")

		if !ok {
			url = "http://" + url
		}

		res, err := http.Get(url)

		if err != nil {
			fmt.Printf("fetch: %v", err)
			os.Exit(1)
		}

		//body, err := ioutil.ReadAll(res.Body)
		// 使用`io.Copy`直接将一个可读流Reader写入到一个可写流Writer
		_, err = io.Copy(os.Stdout, res.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s: %v", url, err)
			os.Exit(1)
		}
	}
}

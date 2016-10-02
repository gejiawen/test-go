package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "URL.Path = %q\n", req.URL.Path)
}

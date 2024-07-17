package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
	fmt.Print("Start server and wait for port 8000")
}

package main

import (
	"fmt"
	"net/http"
)

func Sum(a, b int) int {
	return a + b
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test path")
	})

	if err := http.ListenAndServe("0.0.0.0:80", mux); err != nil {
		fmt.Println(err.Error())
	}
}

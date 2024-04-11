package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	if err := http.ListenAndServe("0.0.0.0:80", mux); err != nil {
		fmt.Println(err.Error())
	}
}

package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Start Go Server localhost:8080")

	http.ListenAndServe(":8080", nil)
}
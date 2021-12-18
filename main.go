package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8080", nil)
}

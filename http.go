package main

import (
	"fmt"
	"net/http" // HL
)

func handler(w http.ResponseWriter, r *http.Request) { // HL
	fmt.Println("running handle...")
	fmt.Fprintf(w, "Hello world via HTTP")
}

func main() {
	http.HandleFunc("/", handler) // HL
	http.ListenAndServe("localhost:9999", nil)
}

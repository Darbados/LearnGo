package main

import (
	"fmt"
	//"net/http"
)

func main() {
	elements := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
	}
	target := 3

	if elements[target] {
		fmt.Println("Element is in the set.")
	} else {
		fmt.Println("Element is not in the set.")
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	//	})

	// http.ListenAndServe(":3658", nil)
}

package main

import (
	"fmt"
)

func printOdds() {
	// Used for a simple test against Python in terms of iteration speed.
	const numElements = 1000000
	odds := make([]int, 0)

	for i := 1; i <= numElements; i++ {
		if i%2 != 0 {
			odds = append(odds, i)
		}
	}

	fmt.Printf("Length of odd numbers slice: %d\n", len(odds))
}

func printSomeStuff() {
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
	printOdds()
}

func main() {
	printSomeStuff()
}

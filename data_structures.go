package main

import "fmt"

// Define a struct type
type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	// Declare and initialize an array of integers
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	fmt.Println("Array ", numbers) // Output: [10 20 30 40 50]

	// Declare and initialize a slice of integers
	slice_numbers := []int{10, 20, 30, 40, 50}

	// Append a new element to the slice
	slice_numbers = append(slice_numbers, 60)

	fmt.Println("Slice ", slice_numbers) // Output: [10 20 30 40 50 60]

	// Declare and initialize a map
	ages := map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 22,
	}

	fmt.Println(ages["Alice"]) // Output: 25

	// Create a struct instance
	person := Person{
		Name:   "Alice",
		Age:    25,
		Gender: "Female",
	}

	fmt.Println(person, " from the struct") // Output: {Alice 25 Female}

	// Declare a variable and initialize it with a value
	num := 42

	// Declare a pointer variable and assign the address of num
	numPtr := &num

	fmt.Println("Value of num:", num)        // Output: Value of num: 42
	fmt.Println("Address of num:", &num)     // Output: Address of num: 0x...
	fmt.Println("Value at numPtr:", *numPtr) // Output: Value at numPtr: 42
}

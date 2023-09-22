package main

import (
	"fmt"
	"math/rand"
)

func _printNumber(number int) {
	fmt.Println("A 'random' generated number: ", number)
}

func randomNumber() {
	var num = rand.Intn(10) + 1
	_printNumber(num)
	num = rand.Intn(10) + 1
	_printNumber(num)
}

func main() {
	randomNumber()
}

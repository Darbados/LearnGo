package main

import "fmt"

func main() {
	for x := 0; x < 5; x++ {
		fmt.Println(x)
	}

	y := 0
	for y < 5 {
		fmt.Println(y)
		y++
	}

	z := 0
	for {
		fmt.Println(z)
		z++
		if z >= 5 {
			break
		}
	}
}

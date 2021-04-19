package main

import "fmt"

func main() {

	// for loop structure
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}

	// how to create a while like loop in Go

	// initialize
	j := 10
	// add only condition
	for j <= 20 {
		fmt.Println(j)
		// increments
		j++
	}
}

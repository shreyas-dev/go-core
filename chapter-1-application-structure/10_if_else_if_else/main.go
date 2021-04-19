package main

import "fmt"

func main() {
	price := 100
	if price > 100 {
		fmt.Println("Price too expensive")
	} else if price == 100 {
		fmt.Println("Price is moderate")
	} else {
		fmt.Println("It's ok")
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	// multiple cases , "," acts like OR operator
	case "Go", "golang":
		fmt.Println("Good that you are learning GO")
		// break not needed as indentation will tell GO to add break
	case "Python", "Pyhton 2.7":
		fmt.Println("You sure you want to do that?")

	case "Java", "Java8", "java":
		fmt.Println("Still a better option than the snake")

	default:
		fmt.Println("Best of luck")
	}

	n := 5

	switch {
	case n%2 == 0:
		fmt.Println("N is even")
	default:
		fmt.Println("N is odd")
	}

	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning ")
	case hour < 16:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good night")
	}

}

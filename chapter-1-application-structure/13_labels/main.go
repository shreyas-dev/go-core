package main

import "fmt"

func main() {

	people := [5]string{"Elon", "Dave", "Mark", "Nelson", "Merry"}
	friends := [2]string{"Mark", "Merry"}

	// outer is a label
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")

	// **GOTO STATEMENT **//

	//the following piece of code creates a loop like a for statement does
	i := 0
loop: // label
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//  goto todo //ERROR it's not permitted to jump over the declaration of x
	//  x := 5
	// todo:
	//  fmt.Println("something here")
}

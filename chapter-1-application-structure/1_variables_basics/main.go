package main

import "fmt"

func main() {

	// using var keyword

	var x int = 6

	var s1 string

	s1 = "Shreyas"

	// Using shorthand
	y := 10

	fmt.Printf("X is %d and Y is %d \n", x, y)

	fmt.Println(s1)

	// un-used variables throw compile time errors
	z := 11
	// we use _ to suppress the error, _ is like a black hole.
	_ = z

	// Multiple declarations of variables

	car, cost := "BMW", 2500000

	fmt.Println("Cost of car:", car, " is inr:", cost)

	// You can only use multiple declaration when atleast
	// one LHS has a new variable , in this case year
	car, cost, year := "Audi", 2500000, 2015

	_ = year

	// Go can check the RHS and understand it's type hence var opened bool is not needed
	var opened = false

	// Here file is a new variable
	opened, file := true, "a.txt"

	// Suppress both the variables in one line
	_, _ = opened, file

	// Better readable declaration
	var (
		salary     float64
		first_name string
		age        int
	)

	fmt.Println(salary, first_name, age)

	// multiple variables declaration of same type
	var a, b, c int

	fmt.Println(a, b, c)

	// We use shorthand sign := when we know the initial value
	// other wise we use var keyword for the default value

	var i, j int

	// multiple value assignment -> multiple declaration
	i, j = 8, 5

	fmt.Println("Before swapping i:", i, " j:", j)

	// swapping two numbers
	i, j = j, i

	fmt.Println("After swapping i:", i, " j:", j)

	// Declaring a variable value using an expression
	sum := 5 + 3.2

	_ = sum
}

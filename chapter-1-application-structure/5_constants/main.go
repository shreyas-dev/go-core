package main

import (
	"fmt"
	// "math"
)

func main() {

	// constant has to be initialized , there's no default value of const
	// To declare a constant and give it a name, we use the const keyword
	// Constants need to be initialized when declared
	const days int = 7 // typed constant
	const pi = 3.14    // untyped constant

	const (
		min1 = -500
		// in a grouped constants
		// constants gets the value from previous constant
		min2

		min3
	)

	// error const can't be initialized in runtime
	// const pwr = math.Pow(10,2)

	// this is allowed as len is a inbuilt function
	// and length is initialised during compile time
	const length = len("hello")

	// Untyped constant
	const x = 2

	var i int = x // x changed to int

	var j float64 = x // x changed to float64

	_, _ = i, j

	const r = 5
	// r has a base int const
	var rr = r

	fmt.Printf("%T\n", rr)

	const (
		a = (iota * 2)
		b // 2
		c // 4
		_ // 6 , if you want to skip a value, use _ identifier
		d // 8
		_ // 10
		e // 12
	)

	fmt.Println(a, b, c, d, e)

	// There are ONLY boolean constants, rune constants, integer constants,
	// floating-point constants, complex constants, and string constants.

	// Declaring multiple (grouped) constants
	const (
		aa         = 5   // untyped constant
		bb float64 = 0.1 // typed constant
	)

	const n, m int = 4, 5

	const (
		m1   = -500
		max1 //gets its type and value form the previous constant. It's 500
		max2 //in a grouped constants, a constant repeats the previous one -> 500
	)

	// CONSTANTS RULES

	// 1. You cannot change a constant
	const temp int = 100
	// temp = 50 //compile-time error

	// 2. You cannot initiate a constant at runtime (constants belong to compile-time)
	// const power = math.Pow(2, 3) //error, functions calls belong to runtime

	// 3. You cannot use a variable to initialize a constant
	t := 5
	// error, variables belong to runtime and you cannot initialize a const to runtime values
	// const tc = t

	// 4. You can use a function like len() to initialize a const if it has as argument
	// a constant string literal (not a variable)
	// a string literal is an untyped constant

	const l1 = len("Hello") //ok

	str := "Hello"
	// const l2 = len(str) //error, str is a variable and belongs to runtime

	_, _ = t, str

	// UNTYPED CONSTANTS
	const y float64 = 1.1

	var v1 int = 5
	var v2 float64 = 1.1

	fmt.Println(x * y)
	// => 5.5, No Error because x is untyped and gets its type when its used first time (float64).

	// fmt.Println(v1 * v2)
	// => Error: invalid operation: v1 * v2 (mismatched types int and float64)
	_, _ = v1, v2

	// IOTA
	// iota is number generator for constants which starts from zero
	// and is incremented by 1 automatically.

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3) // => 0 1 2

	const (
		North = iota //by default 0
		East         //omitting type and value means, repeating its type and value so East = iota = 1 (it increments by 1 automatically)
		South        // -> 2
		West         // -> 3
	)

	// Initializing the constants using a step:
	const (
		c11 = iota * 2 // -> 0
		c22            // -> 2
		c33            // -> 4
	)
}

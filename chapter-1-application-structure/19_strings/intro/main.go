package main

import "fmt"

func main() {
	// using ticks ``

	fmt.Println(`Hello "World"!`) // valid , this is called raw string, interpreted literally

	fmt.Println("Hello\nWorld")

	fmt.Println(`
Hello
World	
	`) // Is same

	// String concatenation
	var str string = "I love" + " Go" + " Programming"
	fmt.Println(str + "!")

	// using backslashes inside a string:
	fmt.Println(`C:\Users\Shreyas`)
	fmt.Println("C:\\Users\\Shreyas")

	// getting an element (byte) of a string:
	fmt.Println("Element at index zero:", s3[0]) // => 73 (ascii code for I)

	fmt.Printf("%s\n", str) // %s for string
	fmt.Printf("%q\n", str) // %q for quoted string

	// strings are immutable and can't be changed
	// s3[5] = 'x' // => error: Cannon assign to s3[5].
}

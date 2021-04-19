package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Please provide two args")
		os.Exit(1)
	}

	fmt.Println("OS args:", os.Args)
	fmt.Println("Path args:", os.Args[0])
	fmt.Println("First Arg:", os.Args[1])
	fmt.Println("No of Args:", len(os.Args))

	var aF, err = strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f\n", aF)
	}

	// the first expression is scoped and the
	// second expression returns it
	if i, err := strconv.Atoi(os.Args[2]); err == nil {
		fmt.Println("No error, i is :", i)
	} else {
		fmt.Println(err)
	}

	// i and err are variables scoped to the if statement only
	if i, err := strconv.Atoi("34"); err == nil {
		fmt.Println("No error. i is ", i)
	} else {
		fmt.Println(err)
	}
}

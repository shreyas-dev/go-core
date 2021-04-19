package main

import "fmt"

func main() {

	// defining your type speed with base of uint
	// uint can be int64 or int32 based on the machine

	// in MacOs and win64 it'll be 64
	type speed uint

	var s1 speed = 32
	var s2 speed = 40

	fmt.Println(s2 - s1)

	// converting speed type to uint
	x := uint(s1)

	var s3 speed = speed(x)

	fmt.Println(s3)

	type km float64
	type mile float64

	var parisToLondon km = 465
	// since mile and km are of same base type
	var distanceInMiles mile

	// you can convert km type to mile by directly calling it
	// and vice versa
	distanceInMiles = mile(parisToLondon) * 0.621

	fmt.Println(distanceInMiles)

	// Aliases,
	// THey don't create a new type, rather it acts like another name
	// to type

	// second is a alias for unit
	type second = uint

	var hour second = 3600

	fmt.Printf("hour type: %T\n", hour) // => hour type: uint

	//no need to convert operations (same type)
	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	// arrays, slices and strings are sliceable
	// slicing doesn't modify the array or the slice, it returns a new one

	// declaring an [5]int array
	a := [5]int{1, 2, 3, 4, 5}

	// a slice expression is formed by specifying a start or a low bound and a stop or high bound like  a[start:stop].
	// this selects a range of elements which includes the element at index start, but excludes the element at index stop.

	// slicing an array returns a slice, not an array
	b := a[1:3]                                 // 1 is called start (included), 3 is called stop (excluded)
	fmt.Printf("Type: %T , Value: %#v\n", b, b) // => Type: []int , Value: []int{2, 3}

	// declaring a slice
	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:3]   //start included, stop excluded
	fmt.Println(s2) //[2 3]

	//for convenience, any of the indexis may be omitted.
	// a missing low index defaults to zero; a missing high index defaults to the length of the sliced operand.
	fmt.Println(s1[2:])       // => [3 4 5 6], same as s1[2:len(s1)]
	fmt.Println(s1[:3])       // => [1 2 3], same as s1[0:3]
	fmt.Println(s1[:])        // => [1 2 3 4 5 6], same with s1[0:len(s1)]
	fmt.Println(s1[:len(s1)]) // => => [1 2 3 4 5 6], returns the entire slice
	// fmt.Println(s1[:45])   //panic: runtime error: slice bounds out of range

	s1 = append(s1[:4], 100) // adds 100 after index 4 (excluded)
	fmt.Println(s1)          // -> [1 2 3 4 100]

	s1 = append(s1[:4], 200) // overwrites the last element
	fmt.Println(s1)          // -> [1 2 3 4 200]

	// a slice expression doesn't create a new backing array. The original and the returned slice are connected!
	s1a := []int{10, 20, 30, 40, 50}
	s3, s4 := s1a[0:2], s1a[1:3] //s3, s4 share the same backing array with s1

	s3[1] = 600      // modifying the backing array so s1, s3 and s4 are in fact modified!!
	fmt.Println(s1a) // -> [10 600 30 40 50]
	fmt.Println(s4)  // -> [600 30]

	// when a slice is created by slicing an array, that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	arr1[1] = 2                       // modifying the array , also modifies slice1 & slice2 as backing array is same
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// append() function creates a complete new slice from an existing slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	// newCars doesn't share the same backing array with cars
	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"                              // only cars is modified
	fmt.Println("cars:", cars, "newCars:", newCars) // => cars: [Nissan Honda Audi Range Rover] newCars: [Ford Honda

	anArray := [5]int{1, 2, 3, 4, 5}
	anSlice := []int{1, 2, 3, 4, 5}

	fmt.Printf("Size of an Array is %d\n", unsafe.Sizeof(anArray)) // Array size is 40 bytes, 8 for each
	fmt.Printf("Size of an slice is %d\n", unsafe.Sizeof(anSlice)) // Slice size is 24 bytes,

	numbers := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers = append(numbers, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n := []int{100, 200, 300}
	numbers = append(numbers, n...) // ... is the ellipsis operator
	fmt.Println(numbers)            // -> [2 3 10 20 30 40 100 200 300]

	//** Slice's Length and Capacity **//

	nums := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 1, Capacity: 1

	nums = append(nums, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums = append(nums, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn) // => [10 20 30] [10 20 30] 3

	if len(os.Args) > 1 {
		var sum, product int64 = 0, 1
		var args = os.Args[1:]
		for _, value := range args {
			num, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				sum += num
				product *= num
			}

		}
		fmt.Printf("Sum: %d , Product: %d\n", sum, product)
	}

}

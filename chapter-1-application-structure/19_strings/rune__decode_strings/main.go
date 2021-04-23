package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str string = "ţarà" // ţară means country in Romanian
	// 't', 'a' ,'r' and 'a' are runes and each rune occupies beetween 1 and 4 bytes.

	//The len() built-in function returns the no. of bytes not runes or chars.
	fmt.Println(len(str)) // -> 6,  4 runes in the string but the length is 6

	// returning the number of runes in the string
	m := utf8.RuneCountInString(str)
	fmt.Println(m) // => 4

	// by using indexes we get the byte at that positioin, not rune.
	fmt.Println("Byte (not rune) at position 1:", str[1]) // -> 163

	// decoding a string byte by byte
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i]) // -> Å£arÄ
	}

	// decoding a string rune by rune manually:
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])

		fmt.Printf("%c", r)
		i += size
	}
	fmt.Println()

	// decoding a string rune by rune automatically:
	for i, r := range str { //the first value returned by range is the index of the byte in string where rune starts
		fmt.Printf("%d -> %c\n", i, r) // => ţară
	}
}

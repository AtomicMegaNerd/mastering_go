package main

import (
	"fmt"
)

func main() {
	aString := "Hello World! 你好！"
	fmt.Println("First character", string(aString[0]))

	// Runes
	r := '你'

	fmt.Println("As an int32 value:", r)

	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// Print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

	// Print an existing string as characters
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}

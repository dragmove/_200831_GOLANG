package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	/*
	// var number uint16 = 260
	number := 6
	fmt.Printf("%T", number)

	var bl bool = true
	fmt.Println(number, bl)
	// type, value
	fmt.Printf("Hello %T, %v", 10, 10)

	// int
	fmt.Printf("%b", 99) // base 2
	fmt.Printf("%o", 99) // base 8
	fmt.Printf("%d", 99) // base 10
	fmt.Printf("%x", 99) // base 16

	// float
	fmt.Printf("%e", 9.98) // 10110001114399639.980000e+00
	fmt.Printf("%f", 9.987654321) // 9.987654
	fmt.Printf("%g", 9.987654321) // 9.987654321

	// string
	fmt.Printf("%s", "Hello world") // Hello world
	fmt.Printf("%q", "Hello world") // "Hello world"
	*/

	// Ref: https://www.youtube.com/watch?v=1-bM3lSBDaA
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you were born: ")

	scanner.Scan()
	// input := scanner.Text()
	// fmt.Printf("You typed: %q", input)
	
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be (%d) years old at the end of 2020", 2020 - input)
}
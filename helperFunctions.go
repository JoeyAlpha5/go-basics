package main

import "fmt"

func getUserInput(input string) string {
	// accepting user input
	fmt.Print("Enter your location: ")
	// & is a pointer to the variable
	fmt.Scan(&input)
	// get the memory address of a variable using & (pointer symbol)
	fmt.Print(&input)
	return input
}

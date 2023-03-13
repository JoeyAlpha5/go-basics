package main

import "fmt"

func main() {
	// declaring variables and constants with data types
	const surname = "Chirwa"
	var name string
	var input string
	name = "Jalome"

	// declare arrays with size
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5

	// inline array declaration
	var array2 = [5]int{2, 4, 6, 8, 10}

	// slice declaration
	var slice []int
	slice = append(slice, 1)

	// concatenating strings, you can also use fmt.Sprintf()
	var concat = name + " " + surname
	fmt.Println(concat)
	fmt.Printf("Hello %s %s \n", name, surname)

	// print the variable type using %T
	fmt.Printf("Type of name is %T \n", name)

	// print the input
	input = getUserInput(input)
	fmt.Printf("\nYour location is %v \n", input)

	// print the array in a loop
	for i := 0; i < len(array); i++ {
		var printOut = fmt.Sprintf("array 1 value: %v, array 2 value %v", array[i], array2[i])
		fmt.Println(printOut)
	}

}

package main

import "fmt"

func main(){
	//formatting strings in go
	// %v is used to print the value of the variable
	// %T is used to print the type of the variable
	// %t is used to print the boolean value
	// %d is used to print the integer value
	// %f is used to print the float value
	// %s is used to print the string value
	// %x is used to print the hexadecimal value
	// %p is used to print the pointer value
	// %e is used to print the scientific notation
	// %q is used to print the double-quoted string
	// %c is used to print the character
	// %b is used to print the binary value
	// %o is used to print the octal value
	// %U is used to print the unicode format
	// %v is used to print the value in default format
	// %q is used to print the double-quoted string

	var name = "John"
	var age = 25
	var isMarried = false
	var height = 5.8
	var weight = 70
	fmt.Printf("Name: %v\n", name)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Is married: %v\n", isMarried)
	fmt.Printf("Height: %v\n", height)
	fmt.Printf("Weight: %v\n", weight)
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of isMarried: %T\n", isMarried)
	fmt.Printf("Type of height: %T\n", height)
	fmt.Printf("Type of weight: %T\n", weight)
	fmt.Printf("Boolean value: %t\n", isMarried)
}
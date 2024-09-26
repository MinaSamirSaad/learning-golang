package main

import "fmt"

// global scope variables
var globalVariable = "Global Variable"

func main() {
	// function scope variables
	// data types goes after the variable name
	// variables can be declared without a value
	// variables have "nil" value if not initialized
	// variables can be declared and initialized in one line
	// string variables declared with double quotes
	var name string
	var age int

	// constants are declared with "const" keyword
	// constants can be declared with or without a type
	// constants can be only boolean, numeric, or string values
	const pi = 3.14

	// variables initialized shortcut
	text := "Hello World"

	name = "John Doe"
	age = 25

	fmt.Println(name, age, pi, text, globalVariable)
	fmt.Println(age)
	fmt.Println(pi)
	fmt.Println(text)
}

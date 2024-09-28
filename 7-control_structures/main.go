package main

import "fmt"

func main() {
	// if statement can put more than one expression but the last one must be a condition
	if a := 1; a == 1 {
		// a is only available in the if statement and else statement
		fmt.Println("a is 1")
	} else {
		fmt.Println("a is not 1")
	}
	// switch statement can take an expression
	switch a := 4; a {
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is not 1 or 2")
		fmt.Println("a is", a)
		// fallthrough keyword can be used to execute the next case statement
		fallthrough
	case 3:
		fmt.Println("a is 3")
	}

	// you can use switch statement without an expression
	// the case statement can take conditions
	// it is like a lot of if else statements
	a := 4
	switch {
	case a == 1:
		fmt.Println("a is 1")
	case a == 2:
		fmt.Println("a is 2")
	default:
		fmt.Println("a is not 1 or 2")
		fmt.Println("a is", a)
	}

	// for statement can be used to iterate over a collection
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for statement can be used to iterate over a collection
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	// for statement can be used to iterate over a map
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// for statement can be used to iterate over a string
	for i, v := range "hello" {
		fmt.Println(i, string(v))
	}
	// there is no while or do while statement in Go
	// but you can use for statement to simulate it
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

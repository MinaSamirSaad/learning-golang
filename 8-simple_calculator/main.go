package main

import (
	"fmt"
	"strings"
)

func main() {
	var operator string
	var a, b int
	continueCalc := "yes"
	fmt.Println("this is a simple calculator")
	for continueCalc == "yes" {
		fmt.Println("which operation do you want to do? (add, subtract, multiply, divide)")
		fmt.Scanf("%s", &operator)
		if !(operator == "add" || operator == "subtract" || operator == "multiply" || operator == "divide") {
			fmt.Println("invalid operator, please try again wth true operator")
			continue
		}
		fmt.Println("enter the first number")
		fmt.Scanf("%d", &a)
		fmt.Println("enter the second number")
		fmt.Scanf("%d", &b)
		switch operator {
		case "add":
			fmt.Println(a, "+", b, "=", a+b)
		case "subtract":
			fmt.Println(a, "-", b, "=", a-b)
		case "multiply":
			fmt.Println(a, "*", b, "=", a*b)
		case "divide":
			fmt.Println(a, "/", b, "=", a/b)
		default:
			fmt.Println("invalid operator")
		}
		fmt.Println("do you want another operator? yes or no")
		fmt.Scanf("%s", &continueCalc)
		// convert exitCalc to lowercase to make the comparison case-insensitive
		continueCalc = strings.ToLower(continueCalc)
		if continueCalc == "no" {
			fmt.Println("goodbye")
		}
	}

}

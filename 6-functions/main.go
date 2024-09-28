package main

import "fmt"

func main() {
	// defer is used to delay the execution of a function until the surrounding function returns
	defer fmt.Println("This is the end")
	a := 1
	fmt.Println(a)
	increment(&a)
	fmt.Println(a)
	for i := a; i > -2; i++ {
		if i >= 20 {
			// panic is used to stop the execution of the program
			panic("This is a panic")
		}
		increment(&a)
	}

	b := "hello"
	fmt.Println(b)
	changeText(&b)
	fmt.Println(b)

	fmt.Println(add(1, 2))

	var slice [5]int
	addItem(&slice)
	addItem(&slice)
	addItem(&slice)
	addItem(&slice)
	fmt.Println(slice)
}

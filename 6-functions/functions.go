package main

func add(a int, b int) int {
	return a + b
}

var item = 0

func addItem(a *[5]int) {
	(*a)[item] = item
	item++
}

// using pointers
func increment(a *int) {
	*a++
}

func changeText(a *string) {
	*a = "hello world"
}

package main

import (
	"fmt"

	"test.com/go/collections/collections"
	"test.com/go/collections/data"
)

func main() {
	printHello()
	fmt.Println(data.PublicVariable)
	fmt.Println(collections.Countries)
	fmt.Println(len(collections.Countries))
	fmt.Println(collections.Slice)
	fmt.Println(collections.Map)
}

package main

import (
	"fmt"
	// this how you import a package from another directory
	// the path is relative to the $moduleName/$packageName
	// in this case, the $moduleName is "test.com/go/packages" you can see it in the go.mod file
	"test.com/go/packages/data"
)

func main() {
	printHello()
	fmt.Println(data.PublicVariable)
}

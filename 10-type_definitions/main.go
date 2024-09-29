package main

import (
	"fmt"

	"test.com/go/types/structs"
)

// Define a new type
type MyInt int
type MyFloat float32

// add methods to the new type
func (mi MyInt) convertToFloat() MyFloat {
	return MyFloat(mi)
}

func (mf MyFloat) convertToInt() MyInt {
	return MyInt(mf)
}

func main() {
	// Use the new type
	var num MyInt = 10
	fmt.Println(num)
	newFloat := num.convertToFloat()
	fmt.Println(newFloat)
	newInt := newFloat.convertToInt()
	fmt.Println(newInt)

	// you can create type aliases
	type MyIntAlias = int
	var numAlias MyIntAlias = 10
	fmt.Println(numAlias)

	// create structs

	user := structs.NewUser(2, "John")
	user.SetId(3)
	fmt.Println(user)
	fmt.Println(user.Print())

	user1 := structs.User{}
	user1.SetId(1)
	user1.SetName("Jane")
	fmt.Println(user1)
}

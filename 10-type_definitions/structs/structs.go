package structs

import "fmt"

// define a struct
// the properties are not exported
// so they can't be accessed from outside the package
// but the methods can be accessed
// the struct is similar to a class in other languages
// the methods are similar to the methods of a class
// the struct can be used outside the package and you can add the data to it
// but you cant access the data directly because it is not exported
// so you can add methods to the struct to access the data like getters and setters
type User struct {
	id   int
	name string
}

// create a new user
// it is factory pattern to create a new user
// you can yse it if you want to add some logic to the creation of the user
// for example you can add a validation to the id or name
// or you use private fields in the struct and you want to set them in the factory
func NewUser(id int, name string) User {
	return User{id: id, name: name}
}

// methods of the User struct like getters and setters
func (u User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

// convert user to string
func (u User) Print() string {
	return fmt.Sprint(u.id) + "----" + u.name
}

// embed a struct
// you can embed a struct in another struct
// the embedded struct is like a property of the struct
// you can access the properties and methods of the embedded struct
// like they are properties and methods of the struct
// you can override the methods of the embedded struct
type Admin struct {
	User
	Role string
}

// create a new admin
func NewAdmin(id int, name string, role string) Admin {
	return Admin{User: User{id: id, name: name}, Role: role}
}

// override the Print method of the User struct
func (a Admin) Print() string {
	return fmt.Sprint(a.id) + "----" + a.name + "----" + a.Role
}

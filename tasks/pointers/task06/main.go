package main

import "fmt"

func main() {
	var user *User = NewUser("Анна", 27)
	(*user).Age = 28
	fmt.Println(user.Name, user.Age)
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

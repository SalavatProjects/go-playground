package main

import "fmt"

func main() {
	var user User = *NewUser("Иван", 27)
	var result1 bool = rename(&user, "Пётр")
	var missingUser *User
	var result2 bool = rename(missingUser, "Анна")
	fmt.Println(result1, result2)
}

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func rename(user *User, newName string) bool {
	if user != nil {
		(*user).Name = newName
		return true
	}

	return false
}

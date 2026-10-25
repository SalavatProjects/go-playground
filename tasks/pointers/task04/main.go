package main

import "fmt"

func main() {
	var user User = User{Name: "Иван", Age: 24}
	user.birthday()
	fmt.Println(user.Age)
}

type User struct {
	Name string
	Age  int
}

func (u *User) birthday() {
	(*u).Age++

}

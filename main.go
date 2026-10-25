package main

import (
	"fmt"
	"go-playground/user"
)

func main() {
	/*var a string = "qwe"

	fmt.Println("Hi", a)

	fmt.Println(calculator.AddFunc(1, 2))
	f := calculator.Calculate(calculator.Subtract)
	fmt.Println(f(1, 2))*/

	var u *user.User = user.NewUser("Иван", 24)
	b := u
	b.Age = 30
	fmt.Println(u.Age)
	fmt.Println(&u)
	fmt.Println(&b)
}

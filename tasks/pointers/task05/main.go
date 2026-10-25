package main

import "fmt"

func main() {
	var original *User = &User{Name: "Иван", Age: 24}
	var sameUser *User = original
	var copiedUser User = *original
	(*sameUser).Age = 30
	copiedUser.Age = 40
	fmt.Println(original.Age)
	fmt.Println(sameUser.Age)
	fmt.Println(copiedUser.Age)

}

type User struct {
	Name string
	Age  int
}

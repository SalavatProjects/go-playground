package main

import "fmt"

func main() {
	var number int = 10
	fmt.Println(number)
	fmt.Println(&number)
	var numberPtr *int = &number
	fmt.Println(*numberPtr)
	*numberPtr = 25
	fmt.Println(number)
}

package main

import (
	"fmt"
	"go-playground/calculator"
)

func main() {
	var a string = "qwe"

	fmt.Println("Hi", a)

	fmt.Println(calculator.AddFunc(1, 2))
	f := calculator.Calculate(calculator.Subtract)
	fmt.Println(f(1, 2))
}

package main

import "fmt"

func main() {
	var firstNumber int = 10
	var secondNumber int = 20
	swap(&firstNumber, &secondNumber)
	fmt.Println(firstNumber, secondNumber)
}

func swap(first *int, second *int) {
	*first, *second = *second, *first
}

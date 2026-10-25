package main

import "fmt"

func main() {
	var number int = 10
	fmt.Println(number)
	increment(&number)
	fmt.Println(number)
}

func increment(number *int) {
	*number++
}

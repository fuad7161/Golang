package main

import "fmt"

func main() {
	var myString = "Hello World"
	fmt.Println(myString)
	changeUsingPointer(&myString)
	fmt.Println(myString)
}

func changeUsingPointer(s *string) {
	newValue := "Hi from Fuad"
	*s = newValue
}

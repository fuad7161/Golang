package main

import "fmt"

func main() {
	mySlice := []string{"dog", "cat", "horse", "banana"}
	fmt.Println(mySlice)

	for _, value := range mySlice {
		fmt.Println(value)
	}
}

package main

import (
	"awesomeProject/handlers"
	"fmt"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomeNumber := handlers.RandomNumber(numPool)
	intChan <- randomeNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)

	num := <-intChan
	fmt.Print(num)
}

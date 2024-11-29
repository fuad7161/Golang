package main

import (
	"errors"
	"log"
)

func main() {
	resut, err := divide(100.0, 0.0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Result of division is: ", resut)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Division by zero not possible")
	}
	result = x / y
	return result, nil
}

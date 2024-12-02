package main

import (
	"errors"
	"fmt"
	"net/http"
)

var portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this a the about page ad 2+2 is %d", sum))
}

// Divide is the handler for two value
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := DivideValues(0.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot Divide by 0")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

// AddValues add two value
func AddValues(x, y int) int {
	return x + y
}

// Divide two value
func DivideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	return x / y, nil
}
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

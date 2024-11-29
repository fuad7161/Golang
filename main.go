package main

import "fmt"

type Animal interface {
	Says() string
	NumberofLegs() int
}

type Dog struct {
	name  string
	Breed string
}

type Gorilla struct {
	name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		name:  "Puppy",
		Breed: "Puppet",
	}
	PringInfo(dog)

	gorilla := Gorilla{
		name:          "Gorilla",
		Color:         "black",
		NumberOfTeeth: 2,
	}

	PringInfo(gorilla)
}

func (g Gorilla) Says() string {
	return "Groooooog"
}
func (g Gorilla) NumberofLegs() int {
	return 2
}
func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberofLegs() int {
	return 4
}

func PringInfo(a Animal) {
	fmt.Println(a.Says(), a.NumberofLegs())
}

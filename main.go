package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	//return &person{name, age}
	p := person{name, age}
	p.age = 43
	return &p
}

func main() {
	fmt.Println(person{"fuad", 20})
	fmt.Println(person{name: "fuad", age: 20})
	fmt.Println(&person{name: "fuad", age: 20})

	Me := struct {
		name   string
		isGood bool
	}{
		"fuadul",
		true,
	}

	anotherme := Me
	fmt.Println(Me)

	anotherme.isGood = false
	fmt.Println(anotherme)
}

package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) showDetails() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func (p *person) setName(name string) {
	p.name = name
}

func (p *person) setAge(age int) {
	p.age = age
}

func (p *person) getName() string {
	return p.name
}
func (p *person) getAge() int {
	return p.age
}
func main() {
	var p person
	p.setName("John")
	p.setAge(18)
	p.showDetails()
}

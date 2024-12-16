package main

import "fmt"

type baseName struct {
	name string
}

func (b *baseName) setName(name string) {
	b.name = name
}

func (b *baseName) getName() string {
	return b.name
}

type persone struct {
	baseName
	age int
}

func (p *persone) setAge(age int) {
	p.age = age
}
func (p *persone) getAge() int {
	return p.age
}

func main() {
	Person := persone{
		baseName: baseName{
			name: "fuad",
		},
		age: 10,
	}
	Person.name = "fuadul"
	Person.baseName.name = "Fuadul Hasan"
	fmt.Println(Person)
	var getset getterSetter
	getset = getterSetter(&Person)
	getset.setAge(12)
	fmt.Println(getset.getName())
	fmt.Println(getset.getAge())

	Name := baseName{
		name: "fuad al hasan",
	}
	fmt.Println(Name)
	// its not possible
	//getset = getterSetter(&Name)

}

type getterSetter interface {
	getName() string
	getAge() int
	setAge(age int)
	setName(name string)
}

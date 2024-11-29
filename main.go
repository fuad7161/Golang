package main

import (
	helper "awesomeProject/handlers"
	"fmt"
)

type myObj struct {
	name string
	age  int
}

func main() {
	var myvar helper.SomeName
	myvar.TypeName = "fuad"
	myvar.TypeNumber = 12
	fmt.Println(myvar)

	var myobj myObj
	myobj.name = "jack"
	myobj.age = 25
	fmt.Println(myobj)
}

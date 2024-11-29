package main

import "fmt"

func main() {
	var firstName, lastName string
	firstName, lastName = saySomething("fuadul", "hasan")
	fmt.Printf("%s %s", firstName, lastName)
}

func saySomething(s, ss string) (string, string) {
	return s, ss
}

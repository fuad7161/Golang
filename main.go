package main

import "fmt"

// obj store in a map

type User struct {
	Name string
	Age  int
}

func main() {
	vis := map[string]User{}

	vis["person1"] = User{
		Name: "Fuadul Hasan",
		Age:  23,
	}

	fmt.Println(vis["person1"].Name)
}

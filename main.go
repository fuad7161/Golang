package main

import (
	"fmt"
	"sort"
)

func main() {

	var mySlice []string
	mySlice = append(mySlice, "fuadul")
	mySlice = append(mySlice, "Hasan")
	fmt.Println(mySlice)
	sort.Strings(mySlice)
	fmt.Println(mySlice)

	var a []int
	a = append(a, 2)
	a = append(a, 1)
	a = append(a, 3)
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println(a)

	b := []int{1, 2, 3, 5, 6, 7, 8, 9}
	fmt.Println(b)
	fmt.Println(b[2:6])
}

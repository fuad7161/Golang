package main

import (
	"fmt"
)

func main() {
	var reverse func(s string) string
	reverse = func(s string) string {
		runes := []rune(s)
		for i := 0; i < len(s)/2; i++ {
			runes[i], runes[len(s)-i-1] = runes[len(s)-i-1], runes[i]
		}
		return string(runes)
	}
	// reverse function
	var Myname func(s string) string
	Myname = func(s string) string {
		fmt.Println("My name is ", s)
		var rs = reverse(s)
		fmt.Println("my name is ", rs)
		return rs
	}

	s := "fuadul"
	fmt.Println(Myname(s))

	// fib recursive function
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(10))
}

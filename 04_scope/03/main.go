package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without clousure, for tow or more funcs to have access to the same variable,
that variable would need to be package scoped

anonymous functions -> assigning a function to a variable
*/

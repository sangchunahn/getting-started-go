package main
// package main where everything starts (you have to have a func call main)

import "fmt"
// fmt = format
// This is how we import in go

func main() {

// you can use a : instead of var
// just like typescript you can declare what the var type is

	message := "Hello Go World!"
	a, b, c := 1, false, 3

	fmt.Printf(message, a, b, c)
}

// Name of a function that is uppercase is exported
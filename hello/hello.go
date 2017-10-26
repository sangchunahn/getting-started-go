package main
// package main where everything starts (you have to have a func call main)

import "fmt"
// fmt = format
// This is how we import in go

func main() {

// you can use a : instead of var
// just like typescript you can declare what the var type is
// star is a pointer

	message := "Hello Go World!"
	var greeting *string = &message
	// greeting is assigned to the memory location of message
	*greeting = "hi"

	fmt.Printf(message, *greeting)
}

// Name of a function that is uppercase is exported
// Pointer is just a special type of variable that holds the address of another variable in memory

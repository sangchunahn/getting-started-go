package main
// package main where everything starts (you have to have a func call main)

import "fmt"
// fmt = format
// This is how we import in go

type Salutation struct {
	name string
	greeting string
}
// custom name for a string

const (
	PI = 3.14
	Language = "Go"
)

const (	
	A = iota
	B = iota
	C = iota
)
// iota is the index location of the const

type Printer func(string) ()

func Greet(salutation Salutation, do Printer) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "Yo")
	do(message)
	do(alternate)
}

// there is a way to ignore the second part of the return. replace it with _

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	fmt.Println(len(greeting))
	// len is the length of whatever and shows you the length
	message = greeting[1] + " " + name
	alternate = "Hey! " + name
	return
}
// ... allows you to utilize multiple parameters
// retrun type is a string
// returns two strings


func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

func main() {

// * Part 1 * // Variable and Pointers

	// you can use a : instead of var
	// just like typescript you can declare what the var type is
	// star is a pointer

	// message := "Hello Go World!"
	// var greeting *string = &message
	// greeting is assigned to the memory location of message
	// *greeting = "hi"

	// fmt.Printf(message, *greeting)


// * Part 2 * // Consts

	// var message Salutation = "Hello World"
	// var s =  Salutation{}
	// s.name = "Bob"

	// fmt.Println(s.name)
	// fmt.Println(s.greeting)


// * Part 3 * // Functions

	// var s =  Salutation{"Bob", "Hello"}
	// Greet(s)

// * Part 3 * // Function Types

	var s =  Salutation{"Bob", "Hello"}
	Greet(s, CreatePrintFunction("!!!"))
}

// Name of a function that is uppercase is exported
// Pointer is just a special type of variable that holds the address of another variable in memory

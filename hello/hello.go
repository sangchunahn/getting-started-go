package main
// package main where everything starts (you have to have a func call main)

import "./greeting"

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

// var s =  greeting.Salutation{"Joe", "Hello"}
// greeting.Greet(s, greeting.CreatePrintFunction("?"), true, 5)
// }

// * Ranges and Slices * //

var s =  greeting.Salutation{"Joe", "Hello"}
greeting.Greet(s, greeting.CreatePrintFunction("?"), true, 5)
}



// Name of a function that is uppercase is exported
// Pointer is just a special type of variable that holds the address of another variable in memory

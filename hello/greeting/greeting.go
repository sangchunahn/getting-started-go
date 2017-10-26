package greeting

import "fmt"
// fmt = format
// This is how we import in go

type Salutation struct {
	Name string
	Greeting string
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

func Greet(salutation Salutation, do Printer, isFormal bool, times int) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)

	for i := 0; i < times; i++ {
		if 	prefix := GetPrefix(salutation.Name); isFormal {
			do(prefix + message)
		} else  {
			do(alternate)
		}
	}
}
// i := 0
// for i < times {
// 	i++
// }
// while loop

// for {
// runs forever
// }

// i := 0
// for {
	// if (i < times) {
	// 	break
	// }
	// if i % 2 {
	// 	continue
	// }
//	i++
// }

// there is a way to ignore the second part of the return. replace it with _

func GetPrefix(name string) (prefix string) {
	switch {
		case name == "Bob": 
			prefix = "Mr "
		case name == "Joe", name == "Amy", len(name) == 10: 
			prefix = "Dr " 
		case name == "Mary":
			prefix = "Mrs " 
		default: 
			prefix = "Dude " 
	}
	return
}
 
func CreateMessage(name string, greeting string) (message string, alternate string) {
	fmt.Println(len(greeting))
	// len is the length of whatever and shows you the length
	message = greeting + " " + name
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

package main

import "fmt"

func main() {
	// use keyword 'var' to declare a variable
	// go is a statically-typed language, so after the name of the variable, you may want to write what the type of it's value is (not always needed)
	// var card string = "Ace of Spades" // this is one way of declaring and initialzing a variable in the same line
	card := "Ace of Spades" // if you use the ':=' operator, go will infer the type of the variable from the right-hand side; effectively the same as declaration on line 8
	// only use ':=' when declaring a new variable; reassignments just use '='
	// variables cannot be ASSIGNED outside of a function body but they can be DECLARED WITHOUT ASSIGNMENT outside of a function body
	fmt.Println(card)
}
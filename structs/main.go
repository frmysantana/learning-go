package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// approach 1 : order of values is assumed to follow top-down order of props in struct definition;
	// author doesn't like this approach since a change in the order of the struct definition
	// would also need to change all of these declarations
	// alex := person{"Alex", "Anderson"}

	// approach 2: explicitly use the key in the declaration
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}

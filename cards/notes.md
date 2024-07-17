# Variable Declarations

-use keyword 'var' to declare a variable
-go is a statically-typed language, so after the name of the variable, you may want to write what the type of it's value is (not always needed)
-`var card string = "Ace of Spades"`
	- this is one way of declaring and initializing a variable in the same line
- `card := "Ace of Spades"` 
	- if you use the ':=' operator, go will infer the type of the variable from the right-hand side; effectively the same as previous declaration
-only use `:=` when declaring a new variable; reassignments just use `=`
-variables cannot be ASSIGNED outside of a function body but they can be DECLARED WITHOUT ASSIGNMENT outside of a function body

# Functions and Return Types

- need to include the return type for a function
- the return type is after the closing parenthesis
- syntax:
```
func functionName() returnType {
	<body_definition>

	return <value_of_type_returnType>
}
```
- Files in the same package can freely call functions defined in other files (in the same package) without any sort of import statement

# Slices and For Loops
- go has 2 built-in types for array-like data structures:
	1. Array - fixed length list of things (like the array type in other compiled languages) - have to declare the size at initialization and STAY AT THAT SIZE
	2. Slice - an array that can grow or shrink (dynamic array like in JS?)
		- syntax: `[]<type_of_record>{recordValue1, recordsValue2,...}`
		- add values by using `append()` function: `append(<slice>, <new_value>)` - this does not mutate `<slice>`, it returns a new slice with the added value
- both array-types need to have the same type for all of their records
- for loops - used to iterate over a closed set
	- syntax: 
	`for <index_variable>, <var_containing_member_of_set> := range <set> { <loop_body> }`
	- range is a keyword that means to iterate over every member of a set
	- every loop of a for loop discards the previous variabes, so the syntax requires a declaration as well as initialization, hence we use `:=`
	- you must use every variable you declare in Go or it will not compile
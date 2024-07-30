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

# OO Approach vs Go Approach
- Go doesn't have classes
- instead of classes, you define custom types made of the more basic types and "functions with a receiver" of that custom type (similar to methods on classes)

# Receiver Functions
- any function that has a reciever included in the declaration syntax:
`func (<receiver_symbol> <receiver_type>) functionName()...`
- doing this gives any variable of type `receiver_type` access to the `functionName` function though a dot syntax similar to using a method from a class
- by convention, the `receiver_symbol` should be a 1 or 2 letter abbreviation of the `receiver_type` - technically you don't have to do this

# Creating a New Deck
- when go syntax requires making a variable that you aren't actually going to use, can use a `_` to tell go to ignore the fact that you aren't using that variable (e.g. for the index variable of a for loop)

# Slice Range Syntax
- go has a slice syntax similar to python: <slice>[startIndexInclusive:endIndexNonInclusive]
- if you leave of `startIndexInclusive`, go will start at 0
- if you leave of `endIndexNonInclusive`, go will include the rest of the slice until the end

# Multiple Return Values
- use a type wrapped in parenthesis after a function's name but before the `{` that starts the definition to denote the return type
- can denote multiple return types if you include the types with comma separators
- e.g. `func name(args) (type1, type2, type3)` denotes that the function returns 3 values with types `type1` `type2` and `type3`
- the return line has each value separated by commas: `return valueOfType1, valueOfType2, valueOfType3`
- can store the returned values in their own variables also by using a comma syntax: `var1, var2, var3 := name(args)`

# Byte Slices
- goal is to save the deck onto the machine - saveToFile
- when looking to do low-level stuff, look at the go docs to see if there is already an implementation
- writeFile is appropriate for this case
- in order to use this function, we need to convert deck (which is a slice of strings) into a slice of bytes (commonly called a `byte slice` in the go world)
- byte slice is a slice of ascii digits that represent the string of your data

# Deck to String
- go can do some type conversion for you: `<type_you_want>(<value_you_currently_have>)`
- so to go from deck (which is a slice of strings) to a byte slice:
	- deck -> one long string -> byte slice

# Joining a Slice of Strings
- import multiple packages with:
```
import (
	"package1"
	"package2"
)
```
- go's string join functionality is `strings.Join()` from the `strings` package

# Saving Data to the Hard Drive
- use the ioutils package for it's WriteFile function

# Reading From the Hard Drive
- use `ReadFile` from ioutils
- may return an error, which is of type error if it has a value or `nil` (this is Go's version of `null`) if there were no errors
- when figuring out how to handle errors in go, ask yourself what would you like the program to do if you were the user?

# Error Handling 
- can split strings with `Split` function from `strings` package

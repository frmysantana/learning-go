package main

// Package means 'project' or 'workspace'; can be thought of as collection of common sourcecode files
// First line of a go file must be the package to which that file belongs to

// Go has 2 types of packages: executable (those that can be run) and re-usable (code used as helpers; used to put re-usable logic)
// Any package named 'main' is automatically treated as executable packages
// Packages with any other name are automatically treated as re-usable packages
// Main packages also need a function malled 'main'

import "fmt" // Import give the file access to all the code in whatever package you're importing
// Import is used to call in re-usable packages as well

func main() {
	fmt.Println("Hi there!")
}

/*
- to run a handful of files: `go run <file1> <file2> ...` - this compiles and runs the files
- to compile: `go build <files>` - this compiles files
- to run compiled files (they are executables): `./path/to/file/<filename>`
- to auto-format all of the code: `go fmt`
- to install packages:
	- `go install` - compiles and installs packages
	- `go get` - downloads raw source code of someone else's package
- to run test files: `go test`
*/

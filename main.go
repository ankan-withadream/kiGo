// First line must be package declaration
package main

// packages inside / outside go
import (
	"fmt"
)

// Global level or package level variable
var name string

// Main function of the program
func main() {
	fmt.Println("Hello world")
	var i int
	name = "ankita"
	fmt.Println(name)

	fmt.Println(i)
	fmt.Print(myfunc())

	// anime variable is init with value "Naruto"
	anime := "Naruto"
	// Passing Reference of the variable anime
	functiontakespointer(&anime)
	// Careful observe the print
	fmt.Println(anime)
}

// Funcation declaration
func myfunc() (string, string) {
	// Short hand variable initialization
	quate := "Hello,"
	return quate, "amit"
}

// Function takes pointers input
func functiontakespointer(s *string) {
	newname := "My hero academia"
	// changing the content of pointer s with the new value which will ultimately change the original value of the referance variable
	*s = newname
}

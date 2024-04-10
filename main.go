// First line must be package declaration
package main

// packages inside / outside go
import (
	"fmt"

	"rsc.io/quote"
)

// Global level or package level variable
var name string

// Main function of the program
func main() {
	fmt.Println(quote.Go())
	var i int
	name = "ankita"
	fmt.Println(name)

	fmt.Println(i)
	fmt.Print(myfunc())
}

// Funcation declaration
func myfunc() (string, string) {
	// Short hand variable initialization
	quate := "Hello,"
	return quate, "amit"
}

// First line must be package declaration
package main

// packages inside / outside go
import (
	"fmt"
	"log"
	"time"
)

// Defining the struct
type User struct {
	Firstname string
	Lastname  string
	Age       int
	DOB       time.Time
}

type Car struct {
	name string
}

// Function declatarion for car struct
func (m *Car) printname() string {
	return m.name
}

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

	// variable newuser created using type User
	newuser := User{
		Firstname: "Amit",
		Lastname:  "Adhikari",
		Age:       22,
		DOB:       time.Now(),
	}

	fmt.Println(newuser)

	newcar := Car{
		name: "AUDI",
	}

	// Calling the function of the struct Car
	log.Println("Care name is:", newcar.printname())

	// Map Data structures
	mymap := make(map[string]string)
	mymap["first"] = "top1"
	mymap["second"] = "top2"
	log.Println(mymap)

	carmap := make(map[int]Car)
	car1 := Car{
		name: "Buggati",
	}
	car2 := Car{
		name: "BMW",
	}
	carmap[1] = car1
	carmap[2] = car2

	log.Println(carmap)

	// Slices, Basically it's array

	var mydogs []string
	mydogs = append(mydogs, "ankan")
	mydogs = append(mydogs, "rana")

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

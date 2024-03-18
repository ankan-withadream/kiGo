package main

import "fmt"

func main3() {
	dog := Dog{}
	cat := Cat{}
	animalSound(dog) // Outputs: Woof!
	animalSound(cat) // Outputs: Meow!
}

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Any type that implements the Animal interface can be passed to this function
func animalSound(a Animal) {
	fmt.Println(a.Speak())
}

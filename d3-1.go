package main

import "fmt"

func main2(khobor string) {
	fmt.Println("Arre !", khobor)

	// show_mfg()
	car1 := Car{wheels: 4, color: "red", isDriving: true}
	bike1 := Bike{wheels: 2, color: "black", isDriving: true, breaks: 2}
	truck1 := Truck{wheels: 10, color: "white", isDriving: false, seats: 3}
	v1 := car1.show_mfg()
	v2 := bike1.show_mfg()
	v3 := truck1.show_mfg()

	fmt.Println(v1, v2, v3)
}


type Car struct {
	wheels int
	color string
	isDriving bool

	// start() bool
	// stop() bool
	// drive() bool
}

type Truck struct {
	wheels int
	color string
	isDriving bool
	seats int
}

type Bike struct {
	wheels int
	color string
	isDriving bool
	breaks int
}

type Vehicle interface {
	show_mfg() string
	if_driving() bool
}


func (c Car) show_mfg() string {
	fmt.Println(c.wheels, c)
	return c.color
}

func (t Truck) show_mfg() int {
	fmt.Println(t)
	return t.seats
}

func (b Bike) show_mfg() int {
	fmt.Println(b)
	return b.breaks
}
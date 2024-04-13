package mutex

// Example of Race condiition

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updatemsg(s string) {
	msg = s
	wg.Done()
}

func Racecondition() {

	msg = "Hello world"

	wg.Add(2)
	go updatemsg("Hello GO")
	go updatemsg("Hello Typescript")
	wg.Wait()
	fmt.Print(msg)
}

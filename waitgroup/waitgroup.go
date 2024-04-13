package waitgroup

import (
	"fmt"
	"sync"
)

func printSome(s string, wg *sync.WaitGroup) {
	fmt.Println(s)
	defer wg.Done()
}

// A goroutine is a lightweight thread managed by the Go runtime
func Waitgroup() {

	// A Wait Group using Go "sync" std library
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"phi",
		"che",
		"zeta",
		"theta",
	}
	// Adding length for that time wait group has to wait
	wg.Add(len(words))
	for i, word := range words {
		// Init a thred by using "go" keyword infront of the command

		go printSome(fmt.Sprintf("%s is at %d", word, i), &wg)
	}

	wg.Wait()

	// Introducting a 1 sec sleep just to wait for returning the result from the threads
	// time.Sleep(1 * time.Second) This is a bad approach to waiting for 1sec

	fmt.Printf("amamit")
	// Observe the output. This is due to not managing the threads we created in the go

}

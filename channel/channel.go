package channel

import (
	"fmt"
	"sync"
)

func Channel() {
	mych := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-mych
		if isChannelOpen {

			fmt.Printf("%d is channel value", val)
		}
		wg.Done()
	}(mych, &wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {

		mych <- 5
		wg.Done()

		close(mych)
		fmt.Printf("%d is channel value", ch)
	}(mych, &wg)

	wg.Wait()
}

package mutex

import (
	"fmt"
	"sync"
)

var waitgroup sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func Mutex() {

	// Bank Balance
	var bankBalance int
	var balance sync.Mutex
	fmt.Printf("Initial Bank balance:%d", bankBalance)
	fmt.Println()

	incomes := []Income{
		{
			Source: "Developer",
			Amount: 25,
		},
		{
			Source: "Blogging",
			Amount: 15,
		},
		{
			Source: "Invesment",
			Amount: 10,
		},
	}

	waitgroup.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer waitgroup.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount

				bankBalance = temp
				balance.Unlock()
				fmt.Printf("on week %d, you earned %d from %s\n", week, income.Amount, income.Source)

			}
		}(i, income)
	}

	waitgroup.Wait()

	// Final Balance

	fmt.Printf("Final Bank balance: %d", bankBalance)
}

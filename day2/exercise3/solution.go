package main

import (
	"fmt"
	"math/rand"
	"sync"
	//"time"
)

func main() {
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	balance := 500

	for debit := 0; debit < 10; debit++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			amount :=rand.Intn(500)

			mutex.Lock()
			fmt.Println("amount debit : ",amount)
			if balance>=amount{
				balance -= amount
			}
			fmt.Println("balance",balance)
			mutex.Unlock()

		}()
	}

	for credit := 0; credit < 10; credit++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			amount :=rand.Intn(500)
			mutex.Lock()
			fmt.Println("amount credit : ",amount)
			balance += amount
			fmt.Println("balance",balance)
			mutex.Unlock()

		}()
	}

	//time.Sleep(5*time.Second)
	wg.Wait()
}


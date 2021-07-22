package main

import (
	"fmt"
	"sync"
)

func main() {
	words := [5]string {"quick","brown","fox","lazy","dog"}
	var state = make(map[string]int)
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup

	for _,word := range words{
		wg.Add(1)
		go func(word string) {
			fmt.Println("in routine for ",word)
			defer wg.Done()
			for _,char := range word{
				mutex.Lock()
				state[string(char)]++
				mutex.Unlock()
			}
		}(word)
	}
	//time.Sleep(5*time.Second)
	wg.Wait()
	fmt.Println(state)
}
